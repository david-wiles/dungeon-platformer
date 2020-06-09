package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	_ "image/png"
	"time"
)

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

func setup() {
	global.gTextures.Load(wTexturePath)

	global.gWorld.Init()
	global.gWorld.LoadMap()

	global.gPlayer.Init(global.gTextures.sprites["gold_knight"], 0, 20 * global.gScale)
	global.gWorld.qt.Insert(global.gPlayer.Bounds)

	global.gCamera.Init()
	global.gCamera.follow = global.gPlayer
	global.gController.Init()
	global.gHud.Init()
	global.gMainMenu.Init()
}

func run() {
	global.gVariables.Load(wConfigFile)

	// Initialize window
	cfg := pixelgl.WindowConfig{
		Title:  wWindowTitle,
		Bounds: pixel.R(float64(global.gVariables.WindowWidth/-2),
			float64(global.gVariables.WindowHeight/-2),
			float64(global.gVariables.WindowWidth/2),
			float64(global.gVariables.WindowHeight/2)),
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

func main() {
	pixelgl.Run(run)
}
