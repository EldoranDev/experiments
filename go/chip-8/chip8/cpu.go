package chip8

import (
	"fmt"
	"github.com/EldoranDev/experiments/go/chip-8/chip8/keyboard"
	"os"

	"github.com/EldoranDev/experiments/go/chip-8/chip8/font"
	"github.com/EldoranDev/experiments/go/chip-8/screen"
	"github.com/EldoranDev/experiments/go/chip-8/stack"
)

const AddressRom = 0x200

const SpriteWidth = 8

type Cpu struct {
	// Flag to indicate an Error state of the Emulator
	Error bool

	// Indicates SuperChip or Chip-48 instruction set
	// SuperChip bool

	// RAM
	Memory [4096]byte

	// Program Counter
	PC uint16

	// index register
	I uint16

	// Address Stack
	Stack *stack.Stack

	// Timer
	Delay byte
	Sound byte

	// General Purpose 16 bit Registers V0 - VF
	V [16]uint8

	// Screen abstraction to interact with game engine
	Screen *screen.Screen

	// Keyboard abstraction to interact with game engine
	Keyboard *keyboard.Keyboard
}

func New() *Cpu {
	return &Cpu{
		Stack:    stack.New(),
		Screen:   screen.New(),
		Keyboard: keyboard.New(),
	}
}

type Op struct {
	Instr uint16

	X uint16
	Y uint16

	N   uint8
	NN  uint8
	NNN uint16
}

func (c *Cpu) Fetch() uint16 {
	var instr uint16
	instr = uint16(c.Memory[c.PC])
	instr <<= 8
	instr |= uint16(c.Memory[c.PC+1])

	return instr
}

func (c *Cpu) Decode(instr uint16) *Op {
	return &Op{
		Instr: instr & 0xF000,

		X:   (instr & 0x0F00) >> 8,
		Y:   (instr & 0x00F0) >> 4,
		N:   uint8(instr & 0x000F),
		NN:  uint8(instr & 0x00FF),
		NNN: instr & 0x0FFF,
	}
}

func (c *Cpu) Execute(instr uint16, op *Op) {
	c.PC += 2

	if instr == 0x00E0 {
		c.Screen.Clear()
		return

	}

	switch op.Instr {
	case 0x0000:
		c.PC, _ = c.Stack.Pop()
		break
	case 0x1000:
		c.PC = op.NNN
		break
	case 0x2000:
		c.Stack.Push(c.PC)
		c.PC = op.NNN
		break
	case 0x3000:
		if c.V[op.X] == op.NN {
			c.PC += 2
		}
		break
	case 0x4000:
		if c.V[op.X] != op.NN {
			c.PC += 2
		}
		break
	case 0x5000:
		if c.V[op.X] == c.V[op.Y] {
			c.PC += 2
		}
		break
	case 0x6000:
		c.V[op.X] = op.NN
		break
	case 0x7000:
		c.V[op.X] += op.NN
		break
	case 0x8000:
		switch op.N {
		case 0x0:
			c.V[op.X] = c.V[op.Y]
			break
		case 0x1:
			c.V[op.X] |= c.V[op.Y]
			break
		case 0x2:
			c.V[op.X] &= c.V[op.Y]
			break
		case 0x3:
			c.V[op.X] ^= c.V[op.Y]
			break
		case 0x4:
			sum16 := uint16(c.V[op.X]) + uint16(c.V[op.Y])

			c.V[op.X] = uint8(sum16)
			if sum16 > 255 {
				c.V[0xF] = 1
			} else {
				c.V[0xF] = 0
			}
			break
		case 0x5:
			if c.V[op.X] > c.V[op.Y] {
				c.V[0xF] = 1
			} else {
				c.V[0xF] = 0
			}

			c.V[op.X] = c.V[op.X] - c.V[op.Y]
		case 0x6:
			if c.V[op.X]%2 == 0 {
				c.V[0xF] = 0
			} else {
				c.V[0xF] = 1
			}

			c.V[op.X] >>= 1
		case 0x7:
			if c.V[op.Y] > c.V[op.X] {
				c.V[0xF] = 1
			} else {
				c.V[0xF] = 0
			}

			c.V[op.X] = c.V[op.Y] - c.V[op.X]
		case 0xE:
			if c.V[op.X] < 128 {
				c.V[0xF] = 0
			} else {
				c.V[0xF] = 1
			}

			c.V[op.X] <<= 1
		default:
			c.Error = true
			fmt.Printf("Unhandled OP 0x8000 Code: %x (%v)\n", op.N, op)
		}

		break
	case 0x9000:
		if c.V[op.X] != c.V[op.Y] {
			c.PC += 2
		}
		break
	case 0xA000:
		c.I = op.NNN
		break
	case 0xD000:
		x := c.V[op.X] & 63 // Wrap X-Coord
		y := c.V[op.Y] & 31 // Wrap Y-Coord

		c.V[0xF] = 0 // Reset Flag

		masks := [8]uint8{128, 64, 32, 16, 8, 4, 2, 1}

		for i := range op.N {
			n := c.I + uint16(i)

			for j := range SpriteWidth {
				set := (c.Memory[n]&masks[j])>>(8-j-1) == 1

				if set {
					c.V[0xF] |= c.Screen.Set(uint16(x)+uint16(j), uint16(y)+uint16(i))
				}
			}
		}
		break
	case 0xE000:
		switch instr & 0x00FF {
		case 0x9E:
			if c.Keyboard.IsPressed(c.V[op.X]) {
				c.PC += 2
			}
			break
		case 0xA1:
			if !c.Keyboard.IsPressed(c.V[op.X]) {
				c.PC += 2
			}
			break
		default:
			fmt.Printf("Unhandled OP 0xE000 Code: %x (%v)\n", op.NN, op)
			c.Error = true
		}
		break
	case 0xF000:
		switch op.NN {
		case 0x1E:
			c.I += uint16(c.V[op.X])

			// Might be needed by the game "Spacefight 2091!"
			if c.I > 0x0FFF {
				c.V[0xF] = 1
			}
			break
		case 0x29:
			c.I = (uint16(c.V[op.X]) * 5) + font.FontAddress
			break
		case 0x33:
			c.Memory[c.I] = c.V[op.X] / 100
			c.Memory[c.I+1] = (c.V[op.X] / 10) % 10
			c.Memory[c.I+2] = c.V[op.X] % 10
			break
		case 0x55:
			for i := uint16(0); i < op.X+1; i++ {
				c.Memory[c.I+i] = c.V[i]
			}

			// TODO: Add Legacy Mode
			// c.I = c.I + op.X + 1
			break
		case 0x65:
			for i := uint16(0); i < op.X+1; i++ {
				c.V[i] = c.Memory[c.I+i]
			}

			// TODO: Add Legacy Mode
			// c.I += op.X + 1
			break
		default:
			fmt.Printf("Unhandled OP 0xF000 Code: %x (%v)\n", op.NN, op)
			break
		}
		break
	default:
		c.Error = true
		fmt.Printf("Unhandled OP Code: %x (%v)\n", op.Instr, op)
	}
}

func (c *Cpu) ReadRom(name string) {
	rom, err := os.ReadFile(fmt.Sprintf("./roms/%s.ch8", name))

	if err != nil {
		panic(err)
	}

	for i, v := range rom {
		c.Memory[i+AddressRom] = v
	}
}

func (c *Cpu) LoadFont(address int) {
	for i, v := range font.Font {
		c.Memory[address+i] = v
	}
}
