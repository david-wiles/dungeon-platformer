package main

import (
	"github.com/faiface/pixel"
)

type Platform struct {
	*CompoundSprite
	*Physics
}

func (p *Platform) Init(s *CompoundSprite, phys *Physics) {
	p.CompoundSprite = s
	p.Physics = phys
	p.Physics.Center = pixel.V(p.Bounds().GetDrawCenter())
}

func (p *Platform) Draw(dt float64) {
	p.CompoundSprite.Draw(p.Physics.Center)
}

func (p *Platform) Center() pixel.Vec {
	return p.Physics.Center
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
