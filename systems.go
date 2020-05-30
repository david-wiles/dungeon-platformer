package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func runRenderSystem(w *World, win *pixelgl.Window) {
	for _, batch := range w.batches {
		batch.Clear()
	}
	for i := range w.entities {
		w.renderStorage[i].draw(pixel.IM.Scaled(pixel.ZV, 4).Moved(w.positionStorage[i]))
	}
	for _, batch := range w.batches {
		batch.Draw(win)
	}
}
