package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	engine, err := NewEngine("Test App", 1280, 720)
	canvas, err := NewCanvas(engine, 0, 0, 1280, 720)

	fmt.Println(err)

	startTime := time.Now()
	totalFrames := 0
	for engine.Running {
		totalFrames++
		if totalFrames >= 100 {
			totalFrames = 0
			startTime = time.Now()
		}
		frames := float64(totalFrames) / time.Now().Sub(startTime).Seconds()

		engine.Clear()
		canvas.Start()
		canvas.SetColor(255, 255, 255, 255)
		vertices := make([]*Vertex, 0)
		for i := 0; i < 1000; i++ {
			vertices = append(vertices, NewVertex(rand.Float32()*1280, rand.Float32()*720, float32(i*10)))
		}
		canvas.DrawVertices(vertices)
		engine.RenderCanvas(canvas)

		engine.Render()
		engine.HandleEvents()
		engine.Window.SetTitle(fmt.Sprintf("%f", frames))


	}

	engine.Cleanup()
}
