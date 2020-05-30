package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math"
)

type Camera struct {
	camPos       pixel.Vec
	camSpeed     float64
	camZoom      float64
	camZoomSpeed float64
}

func (c *Camera) RegisterCameraAction(win *pixelgl.Window, dt float64) {
	if win.Pressed(pixelgl.KeyLeft) {
		c.camPos.X -= c.camSpeed * dt
	}
	if win.Pressed(pixelgl.KeyRight) {
		c.camPos.X += c.camSpeed * dt
	}
	if win.Pressed(pixelgl.KeyDown) {
		c.camPos.Y -= c.camSpeed * dt
	}
	if win.Pressed(pixelgl.KeyUp) {
		c.camPos.Y += c.camSpeed * dt
	}
	c.camZoom *= math.Pow(c.camZoomSpeed, win.MouseScroll().Y)
}

func (c *Camera) Move(win *pixelgl.Window) {
	mat := pixel.IM.Scaled(c.camPos, c.camZoom).Moved(win.Bounds().Center().Sub(c.camPos))
	win.SetMatrix(mat)
}
