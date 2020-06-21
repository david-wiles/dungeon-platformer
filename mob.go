package main

import "github.com/faiface/pixel"

type Mob struct {
	*CompoundSprite
	*Physics
}

// Uses the default Sprite sheet to create a mob
func (m *Mob) Init(s *CompoundSprite, p *Physics) {
	m.CompoundSprite = s
	m.Physics = p
	m.Physics.Center = pixel.V(m.Bounds().GetDrawCenter())
}

func (m *Mob) Draw(dt float64) {
	m.CompoundSprite.Draw(m.Physics.Center)
}

func (m *Mob) Center() pixel.Vec {
	return m.Physics.Center
}

func (m *Mob) Bounds() *Bounds {
	return m.Physics.Bounds
}

func (m *Mob) GetPosition() pixel.Vec {
	return pixel.Vec{m.Bounds().X, m.Bounds().Y}
}

func (m *Mob) Move(move Move) {
	m.Physics.Update(move)
	m.Physics.Center = pixel.V(m.Bounds().GetDrawCenter())
}
