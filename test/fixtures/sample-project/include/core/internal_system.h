#pragma once

#include "core/interfaces.h"
#include <memory>

namespace game_engine {

// Base class that will be inherited by anonymous namespace class
class InternalSystemBase : public Updatable {
public:
    InternalSystemBase() = default;
    virtual ~InternalSystemBase() = default;
    
    virtual void Initialize() = 0;
    virtual void Shutdown() = 0;
};

// Factory function declaration
std::unique_ptr<InternalSystemBase> CreateInternalLogger();

// Test function declaration
void TestAnonymousClasses();

} // namespace game_engine