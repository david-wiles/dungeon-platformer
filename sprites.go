// sprites.go
// This file contains all functions which should be used to get sprites from the sprite sheet
// Since this game only uses one sheet as of now, the batch can be created an passed to each function

package main

import (
	"github.com/faiface/pixel"
	"image"
	"os"
)

var sheet pixel.Picture
var sprites map[string]*pixel.Sprite

// Returns the batch for the dungeon sprite sheet
func getDungeonBatch() *pixel.Batch {
	var err error
	sheet, err = loadPicture("resources/dungeon.png")
	if err != nil {
		panic(err)
	}
	generateSprites(sheet)

	return pixel.NewBatch(&pixel.TrianglesData{}, sheet)
}

// Generate the sprites and store in the global variable in this file
func generateSprites(picture pixel.Picture) {
	sprites = make(map[string]*pixel.Sprite)

	sprites["tall_yellow_creep"] = pixel.NewSprite(sheet, pixel.R(0, 16, 16, 32))
	sprites["short_yellow_creep"] = pixel.NewSprite(sheet, pixel.R(16, 16, 32, 32))
	sprites["tall_green_creep"] = pixel.NewSprite(sheet, pixel.R(32, 16, 48, 32))
	sprites["short_green_creep"] = pixel.NewSprite(sheet, pixel.R(48, 16, 64, 32))

	sprites["short_gray_rock_creep"] = pixel.NewSprite(sheet, pixel.R(0, 32, 16, 48))
	sprites["short_blue_rock_creep"] = pixel.NewSprite(sheet, pixel.R(16, 32, 32, 48))
	sprites["short_flame_creep"] = pixel.NewSprite(sheet, pixel.R(32, 32, 48, 48))
	sprites["short_ghost_creep"] = pixel.NewSprite(sheet, pixel.R(48, 32, 64, 48))

	// The rest of the creeps

	sprites["gold_knight"] = pixel.NewSprite(sheet, pixel.R(240, 0, 256, 16))
	sprites["blue_knight"] = pixel.NewSprite(sheet, pixel.R(224, 0, 240, 16))
	sprites["green_knight"] = pixel.NewSprite(sheet, pixel.R(208, 0, 224, 16))
	sprites["yellow_knight"] = pixel.NewSprite(sheet, pixel.R(192, 0, 208, 16))

	// Some NPCs

	sprites["floor"] = pixel.NewSprite(sheet, pixel.R(0, 240, 16, 256))
	sprites["long_floor"] = pixel.NewSprite(sheet, pixel.R(0, 240, 48, 256))
	sprites["wall"] = pixel.NewSprite(sheet, pixel.R(0, 224, 16, 240))
	sprites["long_wall"] = pixel.NewSprite(sheet, pixel.R(0, 224, 48, 240))
	sprites["slime_wall"] = pixel.NewSprite(sheet, pixel.R(48, 224, 64, 240))
	sprites["low_hole"] = pixel.NewSprite(sheet, pixel.R(48, 240, 64, 256))
	sprites["tall_hole"] = pixel.NewSprite(sheet, pixel.R(64, 240, 80, 256))

	// And the rest
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
