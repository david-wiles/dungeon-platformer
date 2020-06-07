package main

import "github.com/faiface/pixel"

type Mob struct {
	Graphics
	bounds   *Bounds
	hp       float64
	maxLife  float64
	drawFunc func(dt float64)
}

// Uses the default spritesheet to create a mob
func MakeMob(sprite string, x float64, y float64) *Mob {
	var m Mob

	// TODO modify sprite array to store all sprites in global map across multiple batches (graphcs file)
	m = Mob{
		Graphics: Graphics{
			sprite: global.gTextures.sprites[sprite],
			batch:  global.gTextures.batch,
		},
		bounds: &Bounds{
			X:      x,
			Y:      y,
			width:  16,
			height: 16,
			entity: &m,
		},
	}

	return &m
}

func (m *Mob) Init(x, y float64) {

	m.bounds = &Bounds{
		X:      x,
		Y:      y,
		width:  16,
		height: 16,
		entity: m,
	}

}

func (m *Mob) Draw(dt float64) {
	m.sprite.Draw(m.batch, pixel.IM.Scaled(pixel.ZV, global.gScale).Moved(pixel.V(m.bounds.X, m.bounds.Y)))
}

func (m *Mob) GetPosition() pixel.Vec {
	return pixel.Vec{m.bounds.X, m.bounds.Y}
}

func (m *Mob) Move(move pixel.Vec) {
	// TODO handle physics (gravity, etc)
	m.bounds.X += move.X
	m.bounds.Y += move.Y
}
