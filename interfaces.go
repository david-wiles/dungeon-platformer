package main

import "github.com/faiface/pixel"

type Entity interface {
	Init(*Sprite, float64, float64)
	Draw(float64)
	GetPosition() pixel.Vec
	GetDrawVector() pixel.Vec
	Move(pixel.Vec)
}
