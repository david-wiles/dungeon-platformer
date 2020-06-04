package main

import "github.com/faiface/pixel"

type World struct {
	width, height int
	qt            QuadTree
}

func (w *World) Init() {
	w.qt = MakeQuadTree(float64(w.width), float64(w.height))

	batch, sprites := makeDungeon(wResourcePath)
	global.gTextures.batch = batch

	global.gPlayer.Init(100, 100)
	global.gPlayer.batch = batch
	global.gPlayer.sprite = sprites["gold_knight"]
}

// Draw all spites in the current world
func (w *World) Draw(dt float64) {
	// TODO Draw background
	global.gTextures.batch.Clear()

	pos := pixel.ZV
	if global.gCamera.follow != nil {
		// Calculate camera position
		pos = global.gCamera.follow.GetPosition()
		pos.X -= float64(global.gWindowWidth) / 2
		pos.Y -= float64(global.gWindowHeight) / 2
	}

	for _, obj := range w.qt.GetIntersections(&Bounds{X: pos.X, Y: pos.Y, width: float64(global.gWindowWidth), height: float64(global.gWindowHeight)}) {
		obj.entity.Draw(dt)
	}

	global.gTextures.batch.Draw(global.gWin)
}
