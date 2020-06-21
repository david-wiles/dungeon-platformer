package main

import (
	"github.com/faiface/pixel"
)

// Compound sprite contains multiple sprites to draw as if it were one image
// The matrices each correspond to a sprite to describe where it should be drawn relative to
// the center of the entity, represented by the zero vector
type CompoundSprite struct {
	batch   *pixel.Batch
	sprites []struct {
		sprite *pixel.Sprite
		matrix pixel.Matrix
	}
}

// A compound sprite should be initialized by a list of sprites along with the matrix they
// should use to position around a central point
func (s *CompoundSprite) Init(sprites []struct {
	sprite *pixel.Sprite
	matrix pixel.Matrix
}) {
	s.sprites = sprites
}

func (s *CompoundSprite) Append(sprite *pixel.Sprite, matrix pixel.Matrix) {
	s.sprites = append(s.sprites, struct {
		sprite *pixel.Sprite
		matrix pixel.Matrix
	}{
		sprite, matrix,
	})
}

// Draws all sprites using a position vector for the center
func (s *CompoundSprite) Draw(pos pixel.Vec) {
	for _, spr := range s.sprites {
		spr.sprite.Draw(s.batch, spr.matrix.Moved(pos))
	}
}
