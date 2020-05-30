package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"reflect"
)

type World struct {
	userEntityId    uint32
	entities        []uint32
	renderStorage   []RenderComponent
	boxStorage      []BoxComponent
	speedStorage    []SpeedComponent
	positionStorage []pixel.Vec
	batches         map[string]*pixel.Batch
}

func (w *World) Init() {
	w.batches = make(map[string]*pixel.Batch)
}

// This appends all attributes of the entity to each array in the world storage
// Eventually this will need to also sort when adding entities
func (w *World) InsertEntity(e Entity) {
	eRender := e.GetRender()

	if reflect.TypeOf(e) == reflect.TypeOf(PlayerEntity{}) {
		w.userEntityId = e.Id()
	}

	w.entities = append(w.entities, e.Id())
	w.renderStorage = append(w.renderStorage, eRender)
	w.boxStorage = append(w.boxStorage, e.GetBox())
	w.speedStorage = append(w.speedStorage, e.GetSpeed())
	w.positionStorage = append(w.positionStorage, e.GetPosition())

	if w.batches[eRender.batchSource] == nil {
		w.batches[eRender.batchSource] = eRender.batch
	}
}

func (w *World) registerActions(win *pixelgl.Window, dt float64) {
	pos := &w.positionStorage[w.userEntityId]

	if win.Pressed(pixelgl.KeyW) {
		pos.Y += dt * 256
	}
	if win.Pressed(pixelgl.KeyA) {
		pos.X -= dt * 256
	}
	if win.Pressed(pixelgl.KeyS) {
		pos.Y -= dt * 256
	}
	if win.Pressed(pixelgl.KeyD) {
		pos.X += dt * 256
	}
}

func (w *World) RemoveEntity(e Entity) {

}

func (w *World) UpdateSystems() {

}

func (w *World) Tick(win *pixelgl.Window) {
	runRenderSystem(w, win)
}
