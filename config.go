package main

import (
	"encoding/json"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"io/ioutil"
	"os"
)

// Global constants
const (
	wWindowTitle = "Hello from the Dungeon"
	wVersion     = "0.0.1"
	wConfigFile  = "resources/game-config.json"
	wMapFile     = "resources/maps/map-config.json"
	wTexturePath = "resources/textures/dungeon-config.json"
	qtMaxObjects = 8
	qtMaxLevels  = 4
)

// The global Global object
// Allocates memory for each object to be initialized in main
var global = &Global{
	gVsync:      true,
	gWorld:      &World{},
	gTextures:   &Textures{},
	gPlayer:     &Mob{},
	gCamera:     &Camera{},
	gController: &Controller{},
	gWin:        &pixelgl.Window{},
	gClearColor: color.RGBA{0, 0, 0, 1},
	gHud:        &Hud{},
	gMap:        &Map{},
	gMainMenu:   &Menu{},
	gVariables:  &VariableConfig{},
}

// Global struct containing singleton objects for each instance of the game
// gVariables contains global constants that can be modified and saved by the user.
type Global struct {
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

// VariableConfig contains configurations defined by the user.  This includes
// control keys, window size, and other configurations
type VariableConfig struct {
	Vsync        bool    `json:"Vsync"`
	Fullscreen   bool    `json:"Fullscreen"`
	WindowHeight float64 `json:"WindowHeight"`
	WindowWidth  float64 `json:"WindowWidth"`
}

// Load variables config from file
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
