#pragma once

namespace game_engine {
namespace ui {

// UI renderer for 2D interface elements
class Renderer {
public:
    Renderer();
    ~Renderer();
    
    void RenderText(const char* text, int x, int y);
    void RenderButton(const class Button* button);
    void RenderPanel(const class Panel* panel);
    
private:
    class UIContext* ui_context_;
};

} // namespace ui
} // namespace game_engine