package main

import "github.com/faiface/pixel"

type Textures struct {
	batch   *pixel.Batch
	sheet   *pixel.PictureData
	sprites map[string]*pixel.Sprite
	objects []*Sprite
}

type Sprite struct {
	name  string
	pos   pixel.Vec
	scale float64
}

// Loads all textures from specified configuration
func (t *Textures) Load(file string) {

}
