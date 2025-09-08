#pragma once

namespace game_engine {
namespace graphics {

// Low-level graphics renderer for 3D objects
class Renderer {
public:
    Renderer();
    ~Renderer();
    
    void RenderMesh(const class Mesh* mesh);
    void SetShader(const class Shader* shader);
    void Clear();
    
private:
    class RenderContext* context_;
};

} // namespace graphics
} // namespace game_engine