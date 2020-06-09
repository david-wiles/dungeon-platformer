package main

import (
	"github.com/faiface/pixel"
	"math"
)

// Physics is a component of a movable entity
// If the entity should move during the tick, then update should be called
// on the entity's physics object
type Physics struct {
	*Bounds
	Velocity pixel.Vec
	entity   Entity
}

// Update a position based on the current position and move
// Returns the new x, y position of the entity
func (p *Physics) Update(move Move) {
	if move.Jump {
		p.Velocity.Y += 1
	}

	if move.Left {
		p.Velocity.X -= 1
	}

	if move.Right {
		p.Velocity.X += 1
	}

	// Check max velocities

	// Do gravity
	p.Velocity.Y -= 0.2

	// Update position based on velocity and reduce velocity of non-increasing directions
	if math.Abs(p.Velocity.X) > 0 {
		if p.Velocity.X > 0 {
			p.Velocity.X -= math.Abs(p.Velocity.X * 0.1)
		} else {
			p.Velocity.X += math.Abs(p.Velocity.X * 0.1)
		}
	}

	p.X += p.Velocity.X
	p.Y += p.Velocity.Y

	// Calculate collisions
	// TODO change p.Bounds to include the path that p will take during the tick
	intersects := global.gWorld.qt.GetIntersections(p.Bounds)
	if len(intersects) > 1 {
		for _, i := range intersects {
			if i.entity != p.entity {
				var xDiff, yDiff float64
				if p.Velocity.X > 0 {
					xDiff = p.Right() - i.Left()
				} else {
					xDiff = p.Left() - i.Right()
				}

				if p.Velocity.Y > 0 {
					yDiff = p.Top() - i.Bottom()
				} else {
					yDiff = p.Bottom() - i.Top()
				}

				if xDiff > yDiff {
					if p.Velocity.X > 0 {
						p.X = i.Left() - (p.Width * global.gScale)
						p.Velocity.X = 0
					} else {
						p.X = i.Right()
						p.Velocity.X = 0
					}
				} else {
					if p.Velocity.Y > 0 {
						p.Y = i.Bottom() - (p.Height * global.gScale)
						p.Velocity.Y = 0
					} else {
						p.Y = i.Top()
						p.Velocity.Y = 0
					}
				}
			}
		}
	}
}
