package main

import "github.com/faiface/pixel"

// BoxComponent
// Implementation of the hit box for an entity.  This will be used to determine movement
// as well as damage
type BoxComponent struct {
	height, width float32
}

// RenderComponent
// The storage of all data needed to render a specific entity.  This will need to
// include a pointer to the batch as well as a pointer to the specific sprite it
// should use
type RenderComponent struct {
	sprite      *pixel.Sprite
	batch       *pixel.Batch
	batchSource string
}

func (r RenderComponent) draw(m pixel.Matrix) {
	r.sprite.Draw(r.batch, m)
}

// SpeedComponent
// The entity's current speed in x and y direction
type SpeedComponent struct {
	xSpeed float32
	ySpeed float32
}
