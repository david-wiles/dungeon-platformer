package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/go-gl/mathgl/mgl32"
)

type Hud struct {
	canvas         *pixelgl.Canvas
	mapSprite      *pixel.Sprite
	mapFrameCanvas *pixelgl.Canvas
	mapScale       float64
	uPos           mgl32.Vec2
	lifeCanvas     *pixelgl.Canvas
	playerLife     float64
}

func (h *Hud) Init() {

}
