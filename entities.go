package main

import "github.com/faiface/pixel"

type Entity interface {
	Id() uint32
	GetBox() BoxComponent
	GetRender() RenderComponent
	GetSpeed() SpeedComponent
	GetPosition() pixel.Vec
}

type PlayerEntity struct {
	id     uint32
	box    BoxComponent
	render RenderComponent
	speed  SpeedComponent
	pos    pixel.Vec
}

func (p PlayerEntity) Id() uint32 {
	return p.id
}

func (p PlayerEntity) GetBox() BoxComponent       { return p.box }
func (p PlayerEntity) GetRender() RenderComponent { return p.render }
func (p PlayerEntity) GetSpeed() SpeedComponent   { return p.speed }
func (p PlayerEntity) GetPosition() pixel.Vec     { return p.pos }
