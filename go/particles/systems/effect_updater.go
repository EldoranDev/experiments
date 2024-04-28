package systems

import (
	"image/color"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type LifeTimeComponent struct {
	TTL    float32
	MaxTTL float32
}

type effectEntity struct {
	*ecs.BasicEntity
	*LifeTimeComponent
	*common.RenderComponent
}

type EffectUpdaterSystem struct {
	entities map[uint64]effectEntity
	world    *ecs.World
}

func (eu *EffectUpdaterSystem) New(w *ecs.World) {
	eu.world = w
	eu.entities = make(map[uint64]effectEntity)
}

func (eu *EffectUpdaterSystem) Add(basic *ecs.BasicEntity, ttl *LifeTimeComponent, render *common.RenderComponent) {
	if _, ok := eu.entities[basic.ID()]; ok {
		return
	}

	eu.entities[basic.ID()] = effectEntity{basic, ttl, render}
}

func (eu *EffectUpdaterSystem) Update(dt float32) {
	for _, e := range eu.entities {
		e.LifeTimeComponent.TTL -= dt

		e.RenderComponent.Color = color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: uint8(255 * (e.TTL / e.MaxTTL)),
		}

		if e.LifeTimeComponent.TTL <= 0 {
			eu.world.RemoveEntity(*e.BasicEntity)
		}
	}
}

func (eu *EffectUpdaterSystem) Remove(basic ecs.BasicEntity) {
	delete(eu.entities, basic.ID())
}
