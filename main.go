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

		global.gCamera.Update(dt)
		global.gWorld.Draw(dt)

		global.gWin.Update()
	}
}

func setup() {
	global.gWorld.Init()
	global.gCamera.Init()
	global.gTextures.Load(wConfigFile)
	global.gCamera.follow = global.gPlayer
}

func run() {
	// Initialize window
	cfg := pixelgl.WindowConfig{
		Title:  "Hello from pixel!",
		Bounds: pixel.R(0, 0, float64(global.gWindowWidth), float64(global.gWindowHeight)),
		VSync:  true,
	}
	gWin, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	gWin.SetCursorVisible(false)
	global.gWin = gWin

	// Initialize world
	setup()

	gameLoop()
}

func main() {
	pixelgl.Run(run)
}
