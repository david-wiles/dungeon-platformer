package main

import "github.com/faiface/pixel"

type Entity interface {
	Draw(dt float64)
	GetPosition() pixel.Vec
}
