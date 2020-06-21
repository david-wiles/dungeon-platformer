package main

import (
	"encoding/json"
	"github.com/faiface/pixel"
	"os"
)

// A map is a collection of semi-fixed entities in the world.
type Map struct{}

type mapConfig struct {
	MapName  string         `json:"MapName"`
	Entities []entityConfig `json:"Entities"`
}

type entityConfig struct {
	Name   string        `json:"BlockType"`
	X      float64       `json:"X"`
	Y      float64       `json:"Y"`
	Width  float64       `json:"Width"`
	Height float64       `json:"Height"`
	Blocks []blockConfig `json:"Blocks"`
}

// A block config is the configuration for an individual sprite in a compound sprite
// The BlockType should match with a sprite loaded in the textures struct, and each other property
// is a configuration for the matrix used to draw the sprite
type blockConfig struct {
	BlockType string  `json:"BlockType"`
	X         float64 `json:"X"`
	Y         float64 `json:"Y"`
	Rotation  float64 `json:"Rotation"`
	Scale     float64 `json:"Scale"`
}

func (w *World) LoadMap() {
	w.qt.Clear()
	loadMap(wMapFile)
}

func loadMap(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var mConfig mapConfig
	parser := json.NewDecoder(file)
	err = parser.Decode(&mConfig)
	if err != nil {
		panic(err)
	}

	for _, ent := range mConfig.Entities {
		var entity Mob
		cs := &CompoundSprite{
			batch: global.gTextures.batch,
		}

		phys := &Physics{
			Bounds: &Bounds{
				X:      ent.X,
				Y:      ent.Y,
				Width:  ent.Width,
				Height: ent.Height,
				entity: &entity,
			},
			Center:   pixel.ZV,
			Velocity: pixel.ZV,
			entity:   &entity,
		}

		// Append each block to the entity's config
		for _, csConfig := range ent.Blocks {
			cs.Append(
				global.gTextures.sprites[csConfig.BlockType].Frame,
				pixel.IM.Moved(pixel.V(csConfig.X, csConfig.Y)).
					Rotated(pixel.ZV, csConfig.Rotation).
					Scaled(pixel.ZV, csConfig.Scale))
		}

		entity.Init(cs, phys)
		global.gWorld.qt.Insert(entity.Bounds())
	}

}

func (m *Map) Init() {}

// Maybe this will be done with world.draw instead
func (m *Map) drawMap() {}
