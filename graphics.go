package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Graphics struct {
	sheetPath string
	batch     *pixel.Batch
	//batches map[int]*pixel.Batch
	triangles map[int]*pixel.TrianglesData
	sprite    *pixel.Sprite
	canvas    *pixelgl.Canvas
}
