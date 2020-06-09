package main

import (
	"github.com/faiface/pixel"
)

type Platform struct {
	Graphics
	*Bounds
}

func (p *Platform) Init(s *Sprite, x float64, y float64) {
	p.Graphics = Graphics{
		sprite: s.Frame,
		batch:  global.gTextures.batch,
	}
	p.Bounds = &Bounds{
		X:      x,
		Y:      y,
		Width:  s.Width,
		Height: s.Height,
		entity: p,
	}
}

func (p *Platform) Draw(dt float64) {
	p.sprite.Draw(p.batch, pixel.IM.Scaled(pixel.ZV, global.gScale).Moved(p.GetDrawVector()))
}

func (p *Platform) GetPosition() pixel.Vec {
	return pixel.V(p.X, p.Y)
}

func (p *Platform) GetDrawVector() pixel.Vec {
	return pixel.V(p.X+(p.Width*global.gScale)/2, p.Y+(p.Height*global.gScale)/2)
}

func (p *Platform) Move(vec pixel.Vec) {
	// Platforms don't move
}
