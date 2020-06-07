package main

import (
	"github.com/faiface/pixel"
	"math"
)

type Camera struct {
	cam    pixel.Matrix
	zoom   float64
	pos    pixel.Vec
	bounds Bounds
	follow Entity
}

func (c *Camera) Init() {
	c.cam = pixel.IM
	c.zoom = 1
	c.pos = pixel.ZV
	c.bounds = Bounds{
		X:      0,
		Y:      0,
		width:  float64(global.gVariables.WindowWidth),
		height: float64(global.gVariables.WindowHeight),
		entity: nil,
	}
	global.gWin.SetMatrix(c.cam)
}

// Update the camera's position based on the player's position
func (c *Camera) Update(dt float64) {
	pos := c.pos
	if c.follow != nil {
		pos = c.follow.GetPosition()
		pos.X -= float64(global.gVariables.WindowWidth / 2)
		pos.Y -= float64(global.gVariables.WindowHeight / 2)
		c.bounds.Y = pos.Y
		c.bounds.X = pos.X
	}

	// Camera movement
	pos = pixel.Lerp(c.pos, pos, 1-math.Pow(1.0/128, dt))
	c.cam = pixel.IM.Moved(pos.Scaled(-1/c.zoom)).Scaled(pos, c.zoom)
	global.gWin.SetMatrix(c.cam)

	c.pos = pos
}

// Easy access to the bounds object for this camera
func (c *Camera) Bounds() *Bounds {
	return &c.bounds
}
