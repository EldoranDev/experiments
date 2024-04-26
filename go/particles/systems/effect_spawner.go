package systems

import (
	"log"
	"math"
	"math/rand/v2"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type EffectSpawningSystem struct {
	world        *ecs.World
	mouseTracker MouseTracker
}

type MouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}

type Particle struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent

	// Physics Enabled Component
	VelocityComponent
}

func (es *EffectSpawningSystem) New(world *ecs.World) {
	es.world = world

	es.mouseTracker.BasicEntity = ecs.NewBasic()
	es.mouseTracker.MouseComponent = common.MouseComponent{Track: true}

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			sys.Add(&es.mouseTracker.BasicEntity, &es.mouseTracker.MouseComponent, nil, nil)
		}
	}
}

func (*EffectSpawningSystem) Remove(ecs.BasicEntity) {}

func (es *EffectSpawningSystem) Update(dt float32) {
	if engo.Input.Button("SpawnEffect").JustReleased() {

		texture, err := common.LoadedSprite("textures/pixel.png")
		if err != nil {
			log.Fatalln("Unable to load textures: " + err.Error())
			panic("Shutting Down Game")
		}

		for range 1000 {
			p := Particle{BasicEntity: ecs.NewBasic()}

			p.SpaceComponent = common.SpaceComponent{
				Position: engo.Point{X: es.mouseTracker.MouseX, Y: es.mouseTracker.MouseY},
				Width:    2,
				Height:   2,
			}

			p.RenderComponent = common.RenderComponent{
				Drawable: texture,
				Scale:    engo.Point{X: 2, Y: 2},
			}

			theta := rand.Float64() * 2 * math.Pi

			p.VelocityComponent = VelocityComponent{
				X: 300 * float32(math.Cos(theta)) * rand.Float32(),
				Y: 300 * float32(math.Sin(theta)) * rand.Float32(),
			}

			for _, system := range es.world.Systems() {
				switch sys := system.(type) {
				case *common.RenderSystem:
					sys.Add(&p.BasicEntity, &p.RenderComponent, &p.SpaceComponent)
				case *PhysicsSystem:
					sys.Add(&p.BasicEntity, &p.SpaceComponent, &p.VelocityComponent)
				}
			}
		}
	}
}
