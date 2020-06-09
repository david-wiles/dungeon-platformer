package main

import "github.com/faiface/pixel"

type Mob struct {
	Graphics
	Bounds   *Bounds
	hp       float64
	maxLife  float64
	drawFunc func(dt float64)
}

// Uses the default spritesheet to create a mob
func (m *Mob) Init(s *Sprite, x float64, y float64) {
	m.Graphics = Graphics{
		sprite: s.Frame,
		batch:  global.gTextures.batch,
	}
	m.Bounds = &Bounds{
		X:      x,
		Y:      y,
		Width:  s.Width,
		Height: s.Height,
		entity: m,
	}
}

func (m *Mob) Draw(dt float64) {
	m.sprite.Draw(m.batch, pixel.IM.Scaled(pixel.ZV, global.gScale).Moved(m.GetDrawVector()))
}

// Gets the center of the entity
func (m *Mob) GetDrawVector() pixel.Vec {
	return pixel.V(m.Bounds.X+(m.Bounds.Width*global.gScale)/2, m.Bounds.Y+(m.Bounds.Height*global.gScale)/2)
}

func (m *Mob) GetPosition() pixel.Vec {
	return pixel.Vec{m.Bounds.X, m.Bounds.Y}
}

func (m *Mob) Move(move pixel.Vec) {
	// TODO handle physics (gravity, etc)
	m.Bounds.X += move.X
	m.Bounds.Y += move.Y
}
