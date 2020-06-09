package main

import "github.com/faiface/pixel"

type Mob struct {
	Graphics
	*Physics
	DrawCenter pixel.Vec
}

// Uses the default sprite sheet to create a mob
func (m *Mob) Init(s *Sprite, x float64, y float64) {
	m.Graphics = Graphics{
		sprite: s.Frame,
		batch:  global.gTextures.batch,
	}
	m.Physics = &Physics{
		Bounds: &Bounds{
			X:      x,
			Y:      y,
			Width:  s.Width,
			Height: s.Height,
			entity: m,
		},
		Velocity: pixel.ZV,
		entity:   m,
	}
	m.DrawCenter = pixel.V(m.Bounds().GetDrawCenter())
}

func (m *Mob) Draw(dt float64) {
	m.sprite.Draw(m.batch, pixel.IM.Scaled(pixel.ZV, global.gScale).Moved(m.DrawCenter))
}

func (m *Mob) GetDrawCenter() pixel.Vec {
	return m.DrawCenter
}

func (m *Mob) Bounds() *Bounds {
	return m.Physics.Bounds
}

func (m *Mob) GetPosition() pixel.Vec {
	return pixel.Vec{m.Bounds().X, m.Bounds().Y}
}

func (m *Mob) Move(move Move) {
	m.Physics.Update(move)
	m.DrawCenter = pixel.V(m.Bounds().GetDrawCenter())
}
