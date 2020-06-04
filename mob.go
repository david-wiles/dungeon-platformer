package main

import "github.com/faiface/pixel"

type Mob struct {
	Graphics
	bounds   *Bounds
	hp       float64
	maxLife  float64
	drawFunc func(dt float64)
}

func (m *Mob) Init(x, y float64) {

	m.bounds = &Bounds{
		X:      x,
		Y:      y,
		width:  16,
		height: 16,
		entity: m,
	}

	global.gWorld.qt.Insert(m.bounds)

}

func (m *Mob) Draw(dt float64) {
	m.sprite.Draw(m.batch, pixel.IM.Scaled(pixel.ZV, global.gScale).Moved(pixel.V(m.bounds.X, m.bounds.Y)))
}

func (m *Mob) GetPosition() pixel.Vec {
	return pixel.Vec{m.bounds.X, m.bounds.Y}
}
