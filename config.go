package main

import (
	"encoding/json"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"io/ioutil"
	"os"
)

const (
	wWindowTitle = "Hello from the Dungeon"
	wVersion     = "0.0.1"
	wConfigFile  = "resources/game-config.json"
	wMapFile     = "resources/maps/map-config.json"
	wTexturePath = "resources/textures/dungeon-config.json"
	qtMaxObjects = 8
	qtMaxLevels  = 4
)

var global = &Global{
	gScale:    4,
	gVsync:    true,
	gWorld:    &World{},
	gTextures: &Textures{},
	gPlayer: &Mob{
		Graphics: Graphics{},
	},
	gCamera:     &Camera{},
	gController: &Controller{},
	gWin:        &pixelgl.Window{},
	gClearColor: color.RGBA{70, 38, 54, 1},
	gHud:        &Hud{},
	gMap:        &Map{},
	gMainMenu:   &Menu{},
	gVariables:  &VariableConfig{},
}

type Global struct {
	gScale      float64
	gVsync      bool
	gWorld      *World
	gTextures   *Textures
	gPlayer     *Mob
	gCamera     *Camera
	gController *Controller
	gWin        *pixelgl.Window
	gClearColor color.RGBA
	gHud        *Hud
	gMap        *Map
	gMainMenu   *Menu
	gVariables  *VariableConfig
}

type VariableConfig struct {
	Vsync        bool   `json:"Vsync"`
	Fullscreen   bool   `json:"Fullscreen"`
	WindowHeight int    `json:"WindowHeight"`
	WindowWidth  int    `json:"WindowWidth"`
	KeyJump      string `json:"KeyJump"`
	KeyLeft      string `json:"KeyLeft"`
	KeyRight     string `json:"KeyRight"`
	KeyDuck      string `json:"KeyDuck"`
}

// Load config from file
func (v *VariableConfig) Load(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	parser := json.NewDecoder(file)
	err = parser.Decode(v)
	if err != nil {
		panic(err)
	}
}

// Save user preferences
func (v *VariableConfig) Save(filename string) {
	obj, _ := json.Marshal(global.gVariables)
	if err := ioutil.WriteFile(filename, obj, 0644); err != nil {
		panic(err)
	}
}
