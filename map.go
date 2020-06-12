package main

import (
	"encoding/json"
	"os"
)

// A map is a collection of semi-fixed entities in the world.
type Map struct{}

type mapConfig struct {
	MapName  string        `json:"MapName"`
	Entities []blockConfig `json:"Entities"`
}

type blockConfig struct {
	BlockType string  `json:"BlockType"`
	X         float64 `json:"X"`
	Y         float64 `json:"Y"`
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

	for _, obj := range mConfig.Entities {
		var entity Mob
		entity.Init(global.gTextures.sprites[obj.BlockType], obj.X, obj.Y)
		global.gWorld.qt.Insert(entity.Bounds())
	}

}

func (m *Map) Init() {}

// Maybe this will be done with world.draw instead
func (m *Map) drawMap() {}
