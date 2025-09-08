package test

import (
	"testing"
)

func TestHierarchyCommand(t *testing.T) {
	tc := GetTestContext(t)

	t.Run("View hierarchy of Character derived class", func(t *testing.T) {
		result := tc.RunCommand("hierarchy", "Character")
		tc.AssertExitCode(result, 0)
		// Should show inheritance hierarchy
		tc.AssertContains(result.Stdout, "Inherits from:")
		tc.AssertContains(result.Stdout, "└── GameObject")
		tc.AssertContains(result.Stdout, "Character - include/game/character.h")
		tc.AssertContains(result.Stdout, "├── Enemy")
		tc.AssertContains(result.Stdout, "└── Player")
	})

	// Additional test: hierarchy of base class
	t.Run("View hierarchy of GameObject base class", func(t *testing.T) {
		result := tc.RunCommand("hierarchy", "GameObject")
		tc.AssertExitCode(result, 0)
		// Should show what GameObject inherits from
		tc.AssertContains(result.Stdout, "Inherits from:")
		tc.AssertContains(result.Stdout, "└── Updatable")
		// Should show derived classes
		tc.AssertContains(result.Stdout, "GameObject - include/core/game_object.h")
		tc.AssertContains(result.Stdout, "└── Character")
	})

	// Additional test: interface hierarchy
	t.Run("View hierarchy of Updatable interface", func(t *testing.T) {
		result := tc.RunCommand("hierarchy", "Updatable")
		tc.AssertExitCode(result, 0)
		// Should show classes that implement this interface
		tc.AssertContains(result.Stdout, "Updatable - include/core/interfaces.h")
		tc.AssertContains(result.Stdout, "GameObject")
	})

	// Test: multiple matches should show the most relevant one
	t.Run("Hierarchy with multiple class matches", func(t *testing.T) {
		result := tc.RunCommand("hierarchy", "Renderer")
		tc.AssertExitCode(result, 0)
		// Should indicate there are multiple matches (now finds 3: ui::Renderer, graphics::Renderer, MeshRenderer)
		tc.AssertContains(result.Stdout, "matches total, showing most relevant)")
		// Should show the hierarchy of the best match
		tc.AssertContains(result.Stdout, "Renderer - include/")
	})

	// Test: hierarchy with fully qualified name
	t.Run("Hierarchy with namespace-qualified name", func(t *testing.T) {
		result := tc.RunCommand("hierarchy", "game_engine::GameObject")
		tc.AssertExitCode(result, 0)
		// Should find the class even with namespace qualification
		tc.AssertContains(result.Stdout, "GameObject - include/core/game_object.h")
		tc.AssertContains(result.Stdout, "Inherits from:")
		tc.AssertContains(result.Stdout, "└── Character")
	})
}
