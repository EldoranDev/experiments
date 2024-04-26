package main

import (
	"github.com/EldoranDev/experiments/go/particles/config"
	"github.com/EldoranDev/experiments/go/particles/systems"
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type mainScene struct{}

func (*mainScene) Type() string { return "particles" }

func (*mainScene) Preload() {
	engo.Files.Load("textures/pixel.png")
}

func (*mainScene) Setup(u engo.Updater) {
	world, _ := u.(*ecs.World)

	engo.Input.RegisterButton("SpawnEffect", engo.KeySpace)

	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})

	world.AddSystem(&systems.EffectSpawningSystem{})
	world.AddSystem(&systems.PhysicsSystem{})
}

func main() {
	opts := engo.RunOptions{
		Title:  "Particles",
		Height: config.WINDOW_HEIGHT,
		Width:  config.WINDOW_WIDTH,
	}

	engo.Run(opts, &mainScene{})
}
