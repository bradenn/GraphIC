package main

import "github.com/veandco/go-sdl2/sdl"

type Vertex struct {
	point *sdl.FPoint
	vx    float32
	vy    float32
	fx    float32
	fy    float32
	mag   float32
}

func NewVertex(x float32, y float32, m float32) (v *Vertex)  {
	v = &Vertex{
		point: &sdl.FPoint{
			X: x,
			Y: y,
		},
		vx:    0,
		vy:    0,
		fx:    0,
		fy:    0,
		mag:   m,
	}
	return v
}
