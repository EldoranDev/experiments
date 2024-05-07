package main

import (
	"fmt"
	"github.com/EldoranDev/experiments/go/chip-8/chip8"
	"github.com/EldoranDev/experiments/go/chip-8/chip8/font"
	"github.com/EldoranDev/experiments/go/chip-8/screen"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
	"os"
	"time"
)

type Context struct{}

func (ctx *Context) Update() error {
	if cpu.Error {
		return nil
	}

	for key, code := range KeyMap {
		if inpututil.IsKeyJustPressed(key) {
			fmt.Printf("Key Pressed: %v (%x)\n", key, code)
			cpu.Keyboard.PressedKey = code
		}
	}

	cpu.Timer.Tick()

	return nil
}

func (ctx *Context) Draw(img *ebiten.Image) {
	for y := range screen.HEIGHT {
		for x := range screen.WIDTH {
			if cpu.Screen.IsSet(x, y) {
				vector.DrawFilledRect(
					img,
					float32(x*10.),
					float32(y*10.),
					10.,
					10.,
					color.White,
					true,
				)
			}
		}
	}
}

func (ctx *Context) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 320
}

var cpu *chip8.Cpu

func main() {
	ebiten.SetWindowSize(640, 320)

	ebiten.SetTPS(60)

	// Init CPU
	cpu = chip8.New()

	cpu.LoadFont(font.FontAddress)
	cpu.ReadRom(os.Args[1])
	cpu.PC = 0x200

	ticker := time.NewTicker(1000 * time.Microsecond)
	quit := make(chan struct{})

	go func() {
		running := true

		for running && !cpu.Error {
			select {
			case <-ticker.C:
				instr := cpu.Fetch()
				op := cpu.Decode(instr)

				cpu.Execute(instr, op)
				break
			case <-quit:
				running = false
				ticker.Stop()
			}
		}
	}()

	if err := ebiten.RunGame(&Context{}); err != nil {
		log.Fatal(err)
	}

	close(quit)
}
