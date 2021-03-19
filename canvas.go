package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Canvas struct {
	Texture  *sdl.Texture
	Renderer *sdl.Renderer
	Source   *sdl.Rect
	Target   *sdl.Rect
	x        int32
	y        int32
	w        int32
	h        int32
}

func NewCanvas(engine *Engine, x int32, y int32, w int32, h int32) (c *Canvas, err error) {
	c = &Canvas{}
	c.Source = &sdl.Rect{
		X: x,
		Y: y,
		W: w,
		H: h,
	}
	c.Target = c.Source
	c.Renderer = engine.Renderer
	c.Texture, _ = c.Renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, c.w, c.h)

	return c, err
}

func (c *Canvas) Start() {
	_ = c.Renderer.SetRenderTarget(c.Texture)
	c.Renderer.Flush()
}

func (c *Canvas) SetColor(r uint8, g uint8, b uint8, a uint8) {
	_ = c.Renderer.SetDrawColor(r, g, b, a)
}


func (c *Canvas) DrawVertices(vertices []*Vertex) {
	for _, v := range vertices {
		c.Renderer.DrawPointF(v.point.X, v.point.Y)
	}

}
func (c *Canvas) DrawPoint(x float32, y float32) {
	_ = c.Renderer.DrawPointF(x, y)
}

func (c *Canvas) DrawLine(x1 float32, y1 float32, x2 float32, y2 float32) {
	_ = c.Renderer.DrawLineF(x1, y1, x2, y2)
}

func (c *Canvas) GetTexture() *sdl.Texture {
	_ = c.Renderer.SetRenderTarget(nil)
	return c.Texture
}
