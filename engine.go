package main

import "github.com/veandco/go-sdl2/sdl"

type Engine struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer
	Canvases []*Canvas
	Width    int32
	Height   int32
	Running  bool
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func NewEngine(name string, w int, h int) (e *Engine, err error) {
	e = &Engine{Width: int32(w), Height: int32(h)}
	err = sdl.Init(sdl.INIT_EVERYTHING)
	e.Window, err = sdl.CreateWindow(name, -1, -1, int32(w), int32(h), sdl.WINDOW_ALLOW_HIGHDPI)
	e.Renderer, err = sdl.CreateRenderer(e.Window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	_ = e.Renderer.SetScale(2, 2)
	HandleError(err)
	e.Running = true
	return e, nil
}

func (e *Engine) HandleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			e.Running = false
			break
		}
	}
}

func (e *Engine) RenderCanvas(c *Canvas) {
	_ = e.Renderer.Copy(c.GetTexture(), c.Source, c.Target)
}

func (e *Engine) Cleanup() {
	_ = e.Window.Destroy()
	_ = e.Renderer.Destroy()
}

func (e *Engine) Clear() {
	_ = e.Renderer.Clear()
}

func (e *Engine) Render() {
	e.Renderer.Present()
}
