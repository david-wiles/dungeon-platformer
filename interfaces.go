package main

import "github.com/faiface/pixel"

type Entity interface {
	Init(*Sprite, float64, float64)
	Draw(float64)
	GetDrawCenter() pixel.Vec
	Bounds() *Bounds
	GetPosition() pixel.Vec
	Move(Move)
}
