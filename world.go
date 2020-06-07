package main

type World struct {
	width, height int
	qt            QuadTree
}

func (w *World) Init() {
	w.qt = MakeQuadTree(float64(w.width), float64(w.height))
}

// Draw all spites in the current world
func (w *World) Draw(dt float64) {
	global.gTextures.batch.Clear()

	bounds := &Bounds{}
	if global.gCamera.follow != nil {
		bounds = global.gCamera.Bounds()
	}

	for _, obj := range w.qt.GetIntersections(bounds) {
		obj.entity.Draw(dt)
	}

	global.gTextures.batch.Draw(global.gWin)
}
