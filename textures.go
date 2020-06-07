package main

import (
	"encoding/json"
	"github.com/faiface/pixel"
	"image"
	"os"
)

type Textures struct {
	batch   *pixel.Batch
	sprites map[string]*pixel.Sprite
	objects []*Sprite
}

type Sprite struct {
	name  string
	pos   pixel.Vec
	scale float64
}

// For parsing config
type textureConfig struct {
	Filename string         `json:"Filename"`
	Sprites  []spriteConfig `json:"Sprites"`
}

type spriteConfig struct {
	Name string  `json:"Name"`
	X    float64 `json:"X"`
	Y    float64 `json:"Y"`
	W    float64 `json:"W"`
	H    float64 `json:"H"`
}

// Loads all textures from specified configuration
// TODO use json config to loadBatch batches and sprites instead of hardcoded
func (t *Textures) Load(file string) {
	t.batch, t.sprites = loadBatch(file)

	global.gPlayer.batch = global.gTextures.batch
	global.gPlayer.sprite = global.gTextures.sprites["gold_knight"]
}

// Loads a batch and associated sprites from a config file
func loadBatch(path string) (*pixel.Batch, map[string]*pixel.Sprite) {
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
	sprites := make(map[string]*pixel.Sprite)

	for _, s := range tConfig.Sprites {
		sprites[s.Name] = loadSprite(sheet, s)
	}

	return pixel.NewBatch(&pixel.TrianglesData{}, sheet), sprites
}

func loadSprite(p pixel.Picture, s spriteConfig) *pixel.Sprite {
	return pixel.NewSprite(p, pixel.R(s.X, s.Y, s.X+s.W, s.Y+s.H))
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
