package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math"
)

// Camera struct represents the user's view of the world through the current window
// This is used primarily to determine which objects to draw, in order to improve performance
// Follow is a copy of the Entity which the camera is currently centered on
// cam represents the matrix used to set the current window
// zoom is the current scale of the world
// pos is the position of the camera within the world
// bounds represents the boundary of the current window, used to determine which entities 'collide' with
// the window and so should be drawn.
type Camera struct {
	Follow        Entity
	cameraUp      pixelgl.Button
	cameraDown    pixelgl.Button
	cameraLeft    pixelgl.Button
	cameraRight   pixelgl.Button
	cameraLock    pixelgl.Button
	cameraZoomIn  pixelgl.Button
	cameraZoomOut pixelgl.Button
	locked        bool
	cam           pixel.Matrix
	zoom          float64
	pos           pixel.Vec
	bounds        Bounds
}

// Sets all the attributes of the camera to appropriate defaults
func (c *Camera) Init() {
	c.cam = pixel.IM
	c.zoom = 4
	c.pos = pixel.ZV
	c.bounds = Bounds{
		X:      0,
		Y:      0,
		Width:  float64(global.gVariables.WindowWidth),
		Height: float64(global.gVariables.WindowHeight),
		entity: nil,
	}
	c.cameraUp = pixelgl.KeyUp
	c.cameraDown = pixelgl.KeyDown
	c.cameraLeft = pixelgl.KeyLeft
	c.cameraRight = pixelgl.KeyRight
	c.cameraLock = pixelgl.KeySpace
	c.cameraZoomIn = pixelgl.KeyEqual
	c.cameraZoomOut = pixelgl.KeyMinus
	c.locked = true
	global.gWin.SetMatrix(c.cam)
}

// Update the camera's position based on the player's position.
// This should be disabled when the camera is not currently following an entity
func (c *Camera) Update(dt float64) {
	pos := c.pos

	// Check whether to move the camera with follow entity or with user controls
	if c.Follow != nil && c.locked {
		if global.gWin.Pressed(c.cameraLeft) || global.gWin.Pressed(c.cameraRight) ||
			global.gWin.Pressed(c.cameraDown) || global.gWin.Pressed(c.cameraUp) {
			c.locked = false
		}
		pos = c.centerOnFollow()
	} else {
		if global.gWin.Pressed(c.cameraRight) {
			c.pos.X += dt * 250
		}
		if global.gWin.Pressed(c.cameraLeft) {
			c.pos.X -= dt * 250
		}
		if global.gWin.Pressed(c.cameraUp) {
			c.pos.Y += dt * 250
		}
		if global.gWin.Pressed(c.cameraDown) {
			c.pos.Y -= dt * 250
		}
		if global.gWin.Pressed(c.cameraLock) {
			c.locked = true
		}
	}
	if global.gWin.Pressed(c.cameraZoomIn) {
		c.zoom += dt * 5
	}
	if global.gWin.Pressed(c.cameraZoomOut) {
		c.zoom -= dt * 5
	}

	// Camera movement
	c.bounds.X, c.bounds.Y = pos.X, pos.Y
	pos = pixel.Lerp(c.pos, pos, 1-math.Pow(1.0/128, dt))
	c.cam = pixel.IM.Moved(pos.Scaled(-1/c.zoom)).Scaled(pos, c.zoom)
	global.gWin.SetMatrix(c.cam)
	c.pos = pos
}

// Set follow will ensure that the camera's bounds are centered on the entity in
// addition to setting the follow attribute to the given entity.
func (c *Camera) SetFollow(e Entity) {
	c.Follow = e
	c.pos = c.centerOnFollow()
	c.bounds.X, c.bounds.Y = c.pos.X, c.pos.Y
}

func (c *Camera) Bounds() *Bounds { return &c.bounds }

// Returns a vector representing the center of the follow entity
func (c *Camera) centerOnFollow() pixel.Vec {
	pos := c.Follow.GetPosition()
	// Move the camera's position so that the entity's position is in the middle of the screen
	pos.X -= (global.gVariables.WindowWidth*(1/c.zoom))/2 - global.gPlayer.Bounds().Width/2
	pos.Y -= (global.gVariables.WindowHeight*(1/c.zoom))/2 - global.gPlayer.Bounds().Height/2

	return pos
}
