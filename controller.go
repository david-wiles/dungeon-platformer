package main

import (
	"github.com/faiface/pixel/pixelgl"
)

type Controller struct {
	entity   Entity
	keyJump  pixelgl.Button
	keyRight pixelgl.Button
	keyLeft  pixelgl.Button
	keyDuck  pixelgl.Button
}

// Contains possible moves that a user or AI made during the tick
// For example, could contain a move left and a jump
type Move struct {
	Jump  bool
	Duck  bool
	Left  bool
	Right bool
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
	var move Move

	if global.gWin.Pressed(c.keyJump) {
		move.Jump = true
	}
	if global.gWin.Pressed(c.keyLeft) {
		move.Left = true
	}
	if global.gWin.Pressed(c.keyRight) {
		move.Right = true
	}
	if global.gWin.Pressed(c.keyDuck) {
		move.Duck = true
	}

	c.entity.Move(move)
}
