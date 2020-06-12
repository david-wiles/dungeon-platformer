package main

import (
	"encoding/json"
	"github.com/faiface/pixel"
	"image"
	"os"
)

type Textures struct {
	batch   *pixel.Batch
	sprites map[string]*Sprite
}

// Sprite contains a pixel sprite for drawing and the sprite's width and height in pixels,
// which should be used to make the bounds for a rectangular entity
type Sprite struct {
	Frame  *pixel.Sprite
	Width  float64
	Height float64
}

// Texture configuration, from json file
type textureConfig struct {
	Filename string         `json:"Filename"`
	Sprites  []spriteConfig `json:"Sprites"`
}

// Config for an individual sprite
type spriteConfig struct {
	Name string  `json:"Name"`
	X    float64 `json:"X"`
	Y    float64 `json:"Y"`
	W    float64 `json:"W"`
	H    float64 `json:"H"`
}

// Loads all textures from specified configuration
func (t *Textures) Load(file string) {
	t.batch, t.sprites = loadBatch(file)
}

// Loads a batch and associated sprites from a config file
func loadBatch(path string) (*pixel.Batch, map[string]*Sprite) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var tConfig textureConfig
	parser := json.NewDecoder(file)
	err = parser.Decode(&tConfig)
	if err != nil {
		panic(err)
	}

	sheet, err := loadPicture(tConfig.Filename)
	if err != nil {
		panic(err)
	}
	sprites := make(map[string]*Sprite)

	for _, s := range tConfig.Sprites {
		sprites[s.Name] = loadSprite(sheet, s)
	}

	return pixel.NewBatch(&pixel.TrianglesData{}, sheet), sprites
}

func loadSprite(p pixel.Picture, s spriteConfig) *Sprite {
	return &Sprite{
		Frame:  pixel.NewSprite(p, pixel.R(s.X, s.Y, s.X+s.W, s.Y+s.H)),
		Width:  s.W,
		Height: s.H,
	}
}

// Load a picture from an image file
func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}
