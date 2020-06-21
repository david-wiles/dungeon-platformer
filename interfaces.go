package main

import "github.com/faiface/pixel"

type Entity interface {
	Init(*CompoundSprite, *Physics)
	Draw(float64)
	Center() pixel.Vec
	Bounds() *Bounds
	GetPosition() pixel.Vec
	Move(Move)
}
