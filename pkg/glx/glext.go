package glx

import "github.com/o5h/engine/internal/opengl/gl"

type BufferUsage int
type DrawMode uint32

const (
	DynamicDraw BufferUsage = gl.DYNAMIC_DRAW
	StaticDraw  BufferUsage = gl.STATIC_DRAW

	Points        DrawMode = gl.POINTS
	Lines         DrawMode = gl.LINES
	LineLoop      DrawMode = gl.LINE_LOOP
	LineStrip     DrawMode = gl.LINE_STRIP
	Triangles     DrawMode = gl.TRIANGLES
	TriangleStrip DrawMode = gl.TRIANGLE_STRIP
	TriangleFan   DrawMode = gl.TRIANGLE_FAN
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}
