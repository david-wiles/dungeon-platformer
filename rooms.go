package main

import "github.com/faiface/pixel"

func drawMainPlatform(batch *pixel.Batch) {

	// Draw floor
	sprites["long_floor"].Draw(batch, pixel.IM.Scaled(pixel.ZV, 4).Moved(pixel.Vec{-48 * 4, 0}))
	sprites["long_floor"].Draw(batch, pixel.IM.Scaled(pixel.ZV, 4))
	sprites["long_floor"].Draw(batch, pixel.IM.Scaled(pixel.ZV, 4).Moved(pixel.Vec{48 * 4, 0}))

	sprites["long_wall"].Draw(batch, pixel.IM.Scaled(pixel.ZV, 4).Moved(pixel.Vec{-48 * 4, -16 * 4}))
	sprites["long_wall"].Draw(batch, pixel.IM.Scaled(pixel.ZV, 4).Moved(pixel.Vec{0, -16 * 4}))
	sprites["long_wall"].Draw(batch, pixel.IM.Scaled(pixel.ZV, 4).Moved(pixel.Vec{48 * 4, -16 * 4}))

}
