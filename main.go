package main

import (
	"github.com/david-wiles/goplai"
	"github.com/faiface/pixel"
	_ "image/png"
)

func main() {
	goplai.Run(func() {
		goplai.InitGlobal(goplai.Config{
			WindowTitle: "Welcome to the Dungeon",
			VarPath:     "resources/game-config.json",
		})
		goplai.G.Textures.Load("resources/textures/dungeon-config.json")
		goplai.G.World.LoadMap("resources/maps/map-config.json")

		goplai.SetPlayer(goplai.NewMob(goplai.G.Textures.GetSprite("gold_knight"), &goplai.Physics{
			Bounds: &goplai.Bounds{
				X:      48,
				Y:      21,
				Width:  16,
				Height: 16,
			},
			Center:   pixel.ZV,
			Velocity: pixel.ZV,
		}))

		goplai.BeginSimulation()
	})
}
