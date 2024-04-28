package systems

import (
	"github.com/EldoranDev/experiments/go/particles/config"
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type physicsEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*VelocityComponent
}

type VelocityComponent struct {
	X float32
	Y float32
}

type PhysicsSystem struct {
	entities map[uint64]physicsEntity
	world    *ecs.World
}

func (ps *PhysicsSystem) New(w *ecs.World) {
	ps.world = w
	ps.entities = make(map[uint64]physicsEntity)
}

func (ps *PhysicsSystem) Add(basic *ecs.BasicEntity, space *common.SpaceComponent, vel *VelocityComponent) {
	if _, ok := ps.entities[basic.ID()]; ok {
		return
	}

	ps.entities[basic.ID()] = physicsEntity{basic, space, vel}

}

func (ps *PhysicsSystem) Update(dt float32) {
	for _, e := range ps.entities {
		e.VelocityComponent.Y += 9.82

		e.SpaceComponent.Position.X += e.VelocityComponent.X * dt
		e.SpaceComponent.Position.Y += e.VelocityComponent.Y * dt

		if e.SpaceComponent.Position.X < 0 ||
			e.SpaceComponent.Position.Y < 0 ||
			e.SpaceComponent.Position.X >= config.WINDOW_WIDTH+20 ||
			e.SpaceComponent.Position.Y >= config.WINDOW_HEIGHT+20 {

			ps.world.RemoveEntity(*e.BasicEntity)
		}
	}
}

func (ps *PhysicsSystem) Remove(basic ecs.BasicEntity) {
	delete(ps.entities, basic.ID())
}
