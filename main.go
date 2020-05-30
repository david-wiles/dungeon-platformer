package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	_ "image/png"
	"time"
)

func run() {
	// Initialize window
	cfg := pixelgl.WindowConfig{
		Title:  "Hello from pixel!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Initialize world
	path := "resources/dungeon.png"
	batch, sprites := makeDungeon(path)
	world := &World{}
	world.Init()
	world.InsertEntity(&PlayerEntity{
		id:  0,
		box: BoxComponent{16, 16},
		render: RenderComponent{
			sprite:      sprites["gold_knight"],
			batch:       batch,
			batchSource: "resources/dungeon.png",
		},
		speed: SpeedComponent{
			xSpeed: 0,
			ySpeed: 0,
		},
		pos: pixel.Vec{0, 0},
	})

	// Initialize camera
	camera := &Camera{
		camPos:       pixel.ZV,
		camSpeed:     500.0,
		camZoom:      1.0,
		camZoomSpeed: 1.1,
	}

	// Game loop
	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		world.registerActions(win, dt)

		camera.Move(win)
		camera.RegisterCameraAction(win, dt)

		win.Clear(color.RGBA{70, 38, 54, 1})

		world.Tick(win)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
