package main

import "github.com/faiface/pixel/pixelgl"

type Menu struct {
	items            []*MenuItem
	videoModes       []pixelgl.VideoMode
	currentVideoMode int
}

type MenuItem struct {
	action   func()
	nextItem func()
	lastItem func()
	name     string
	canvas   *pixelgl.Canvas
	scale    float64
	selected int
}

func (m *Menu) Init() {}

func (m *Menu) MakeMain() {}

func (m *Menu) addMenuItem(scale float64, name string, action func()) {}

// Add other functions
