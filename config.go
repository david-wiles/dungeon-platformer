package main

import (
	"encoding/json"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"io/ioutil"
	"os"
)

const (
	wConfigFile   = "config.json"
	wResourcePath = "resources/dungeon.png"
	qtMaxObjects  = 8
	qtMaxLevels   = 4
)

var global = &Global{
	gWindowHeight: 768,
	gWindowWidth:  1024,
	gVsync:        true,
	gWorld:        &World{},
	gTextures:     &Textures{},
	gPlayer: &Mob{
		Graphics: Graphics{},
	},
	gCamera:     &Camera{},
	gWin:        &pixelgl.Window{},
	gClearColor: color.RGBA{70, 38, 54, 1},
	gUI:         &Hud{},
	gMap:        &Map{},
	gMainMenu:   &Menu{},
	gVariables:  &VariableConfig{},
	gScale:      4,
}

type Global struct {
	gWindowHeight int
	gWindowWidth  int
	gVsync        bool
	gWorld        *World
	gTextures     *Textures
	gPlayer       *Mob
	gCamera       *Camera
	gWin          *pixelgl.Window
	gClearColor   color.RGBA
	gUI           *Hud
	gMap          *Map
	gMainMenu     *Menu
	gVariables    *VariableConfig
	gScale        float64
}

type VariableConfig struct {
	Vsync        bool `json:"Vsync"`
	Fullscreen   bool `json:"Fullscreen"`
	WindowHeight int  `json:"WindowHeight"`
	WindowWidth  int  `json:"WindowWidth"`
	KeyJump      int  `json:"KeyJump"`
	KeyLeft      int  `json:"KeyLeft"`
	KeyRight     int  `json:"KeyRight"`
}

// Load config from file
func (v *VariableConfig) LoadConfiguration() {
	file, err := os.Open(wConfigFile)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	parser := json.NewDecoder(file)
	err = parser.Decode(v)
	if err != nil {
		panic(err)
	}
}

// Save user preferences
func (v *VariableConfig) SaveConfiguration() {
	json, _ := json.Marshal(global.gVariables)
	if err := ioutil.WriteFile(wConfigFile, json, 0644); err != nil {
		panic(err)
	}
}
