package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Controller struct {
	entity   Entity
	keyJump  pixelgl.Button
	keyRight pixelgl.Button
	keyLeft  pixelgl.Button
	keyDuck  pixelgl.Button
}

// Register action buttons from config file
func (c *Controller) Init() {
	c.entity = global.gPlayer
	c.keyJump = pixelgl.KeyW
	c.keyRight = pixelgl.KeyD
	c.keyLeft = pixelgl.KeyA
	c.keyDuck = pixelgl.KeyS
}

// Listen for keys pressed and updates position of the entity
func (c *Controller) Update(dt float64) {
	var move pixel.Vec

	if global.gWin.Pressed(c.keyJump) {
		move.Y += dt * 500
	}
	if global.gWin.Pressed(c.keyLeft) {
		move.X -= dt * 500
	}
	if global.gWin.Pressed(c.keyRight) {
		move.X += dt * 500
	}
	if global.gWin.Pressed(c.keyDuck) {
		move.Y -= dt * 500
	}

	c.entity.Move(move)
}
