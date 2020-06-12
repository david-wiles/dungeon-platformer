package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	_ "image/png"
	"time"
)

// Game loop which conducts the game. This represents each 'tick' of the world
func gameLoop() {
	last := time.Now()
	for !global.gWin.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		global.gWin.Clear(global.gClearColor)

		// Update systems

		global.gController.Update(dt)
		global.gCamera.Update(dt)
		global.gWorld.Draw(dt)

		global.gWin.Update()
	}
}

// Sets up all global variables and objects
func setup() {
	global.gTextures.Load(wTexturePath)

	global.gWorld.Init()
	global.gWorld.LoadMap()

	global.gPlayer.Init(global.gTextures.sprites["gold_knight"], 48, 21)
	global.gWorld.qt.Insert(global.gPlayer.Bounds)

	global.gCamera.Init()
	global.gCamera.SetFollow(global.gPlayer)

	global.gController.Init()
	global.gHud.Init()
	global.gMainMenu.Init()
}

// Function which runs execution of the program in window
// Everything in this loop will be run on the main thread.
// It is only necessary that graphics code run on this thread
func run() {
	global.gVariables.Load(wConfigFile)

	// Initialize window
	cfg := pixelgl.WindowConfig{
		Title:  wWindowTitle,
		Bounds: pixel.R(0, 0, global.gVariables.WindowWidth, global.gVariables.WindowHeight),
		VSync:  true,
	}
	gWin, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	gWin.SetCursorVisible(false)
	global.gWin = gWin

	setup()

	gameLoop()
}

// Entry point for the program
func main() {
	pixelgl.Run(run)
}
