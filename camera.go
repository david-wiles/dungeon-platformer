package main

import (
	"github.com/faiface/pixel"
)

type Camera struct {
	cam    pixel.Matrix
	zoom   float64
	pos    pixel.Vec
	follow Entity
}

func (c *Camera) Init() {
	c.setPosition(0, 0)
	c.zoom = 1
	c.follow = global.gPlayer
	global.gWin.SetMatrix(c.cam)
}

func (c *Camera) setPosition(x, y float64) {
	c.pos = pixel.V(x, y)
}

// Update the camera's position based on the player's position
func (c *Camera) Update(dt float64) {
	pos := c.pos
	if c.follow != nil {
		pos = c.follow.GetPosition()
	}

	// Camera movement
	//pos = pixel.Lerp(c.pos, pos, 1 - math.Pow(1.0 / 128, dt))
	//c.cam = pixel.IM.Scaled(pos, c.zoom).Moved(pos)
	//global.gWin.SetMatrix(c.cam)
	global.gWin.SetMatrix(pixel.IM.Moved(global.gWin.Bounds().Center()))

	c.pos = pos
}
