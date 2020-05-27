package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	_ "image/png"
	"math"
	"time"
)

func run() {
	// Create window
	cfg := pixelgl.WindowConfig{
		Title:  "Hello from pixel!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	var (
		camPos       = pixel.ZV
		camSpeed     = 500.0
		camZoom      = 1.0
		camZoomSpeed = 1.1
		last         = time.Now()

		playerPos = pixel.ZV
	)

	batch := getDungeonBatch()

	drawMainPlatform(batch)

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()
		cam := pixel.IM.Scaled(camPos, camZoom).Moved(win.Bounds().Center().Sub(camPos))
		win.SetMatrix(cam)

		// Player actions
		if win.Pressed(pixelgl.KeyW) {
			playerPos.Y += dt * 256
		}
		if win.Pressed(pixelgl.KeyA) {
			playerPos.X -= dt * 256
		}
		if win.Pressed(pixelgl.KeyS) {
			playerPos.Y -= dt * 256
		}
		if win.Pressed(pixelgl.KeyD) {
			playerPos.X += dt * 256
		}

		// Camera actions
		if win.Pressed(pixelgl.KeyLeft) {
			camPos.X -= camSpeed * dt
		}
		if win.Pressed(pixelgl.KeyRight) {
			camPos.X += camSpeed * dt
		}
		if win.Pressed(pixelgl.KeyDown) {
			camPos.Y -= camSpeed * dt
		}
		if win.Pressed(pixelgl.KeyUp) {
			camPos.Y += camSpeed * dt
		}
		camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)

		win.Clear(color.RGBA{70, 38, 54, 1})
		batch.Draw(win)
		sprites["gold_knight"].Draw(win, pixel.IM.Scaled(pixel.ZV, 4).Moved(playerPos))
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
