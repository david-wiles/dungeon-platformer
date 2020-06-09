package main

import (
	"github.com/faiface/pixel"
)

type Platform struct {
	Graphics
	*Physics
	DrawCenter pixel.Vec
}

func (p *Platform) Init(s *Sprite, x float64, y float64) {
	p.Graphics = Graphics{
		sprite: s.Frame,
		batch:  global.gTextures.batch,
	}
	p.Physics = &Physics{
		Bounds: &Bounds{
			X:      x,
			Y:      y,
			Width:  s.Width,
			Height: s.Height,
			entity: p,
		},
		Velocity: pixel.ZV,
		entity:   p,
	}
	p.DrawCenter = pixel.V(p.Bounds().GetDrawCenter())
}

func (p *Platform) Draw(dt float64) {
	p.sprite.Draw(p.batch, pixel.IM.Scaled(pixel.ZV, global.gScale).Moved(p.DrawCenter))
}

func (p *Platform) GetDrawCenter() pixel.Vec {
	return p.DrawCenter
}

func (p *Platform) Bounds() *Bounds {
	return p.Physics.Bounds
}

func (p *Platform) GetPosition() pixel.Vec {
	return pixel.V(p.X, p.Y)
}

func (p *Platform) Move(Move) {
	// Platforms don't move ;.;
}
