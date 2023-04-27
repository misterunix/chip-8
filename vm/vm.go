package vm

import (
	"math/rand"
	"os"
	"sync"
	"time"
)

type Chip8 struct {
	MU          sync.Mutex  // Mutex
	Display     []uint8     // Display pixels.
	DisplaySize int         // Display size in bytes
	ScreenSize  [2]int      // x,y
	Memory      [4096]uint8 // 4096 bytes of memory
	Registers   [16]uint8   // V0-VF
	Stack       [16]uint16  // 16 bit stack
	Keys        [16]uint8   // 16 keys
	I           uint16      // Index register
	ST          uint8       // sound timer
	DT          uint8       // delay timer
	PC          uint16      // Program Counter
	SP          uint8       // stack pointer
	OpCode      uint16      // 2 bytes opcode
	Width       uint8       // Screen width
	Height      uint8       // Screen height
	Draw        bool        // Draw flag
	rnd         *rand.Rand  // Random number generator
	x           uint8       // x register
	y           uint8       // y register
	n           uint8       // n nibble
	nn          uint8       // nn byte
	nnn         uint16      // nnn address
	debug       bool        // debug flag
	DebugString string      // debug string
}

func (v *Chip8) init() {
	v.rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func (v *Chip8) clearScreen() {
	v.Draw = true
	for i := 0; i < v.DisplaySize; i++ {
		v.Display[i] = 0
	}
}

func (v *Chip8) floodDisplay() {
	for i := 0; i < v.DisplaySize; i++ {
		v.Display[i] = 1
	}
}

func NewVM(model int) *Chip8 {
	switch model {
	case 0:
		v := Chip8{
			ScreenSize: [2]int{64, 32},
			debug:      true,
			Draw:       false,
		}
		v.DisplaySize = v.ScreenSize[0] * v.ScreenSize[1]
		v.Display = make([]uint8, v.DisplaySize)
		return &v

	case 1:
		v := Chip8{
			ScreenSize: [2]int{128, 64},
			debug:      false,
			Draw:       false,
		}
		v.DisplaySize = int(v.ScreenSize[0]) * int(v.ScreenSize[1])
		v.Display = make([]uint8, v.DisplaySize)

		return &v

	default:
		v := Chip8{
			ScreenSize: [2]int{64, 32},
			debug:      false,
			Draw:       false,
		}
		v.DisplaySize = int(v.ScreenSize[0]) * int(v.ScreenSize[1])
		v.Display = make([]uint8, v.DisplaySize)

		return &v
	}

}

func (v *Chip8) Reset() {
	v.init()
	v.PC = 0x200 // Program Start
	v.DT = 0     // Delay timer
	v.ST = 0     // Sound timer
	v.SP = 0     // Stack pointer
	v.I = 0      // Index register
	for i := 0; i < 16; i++ {
		v.Registers[i] = 0
	}
	//	for i:=0;i<len(v.Memory);i++{
	//v.Memory[i] = 0
	//}
	//v.Display = make([]uint8, (v.Width*v.Height)/8) // Display pixels.
	v.clearScreen()
	//fmt.Println("Reset")
	//v.floodDisplay()
	v.loadFontSet()
}

// Load fonts to the beginning of memory
func (v *Chip8) loadFontSet() {

	// Load fontset from top down

	// 0
	v.Memory[0x00] = 0xF0
	v.Memory[0x01] = 0x90
	v.Memory[0x02] = 0x90
	v.Memory[0x03] = 0x90
	v.Memory[0x04] = 0xF0

	// 1
	v.Memory[0x5] = 0x20
	v.Memory[0x6] = 0x60
	v.Memory[0x7] = 0x20
	v.Memory[0x8] = 0x20
	v.Memory[0x9] = 0x70

	// 2
	v.Memory[0x0A] = 0xF0
	v.Memory[0x0B] = 0x10
	v.Memory[0x0C] = 0xF0
	v.Memory[0x0D] = 0x80
	v.Memory[0x0E] = 0xF0

	// 3
	v.Memory[0x0F] = 0xF0
	v.Memory[0x10] = 0x10
	v.Memory[0x11] = 0xF0
	v.Memory[0x12] = 0x10
	v.Memory[0x13] = 0xF0

	// 4
	v.Memory[0x14] = 0x90
	v.Memory[0x15] = 0x90
	v.Memory[0x16] = 0xF0
	v.Memory[0x17] = 0x10
	v.Memory[0x18] = 0x10

	// 5
	v.Memory[0x19] = 0xF0
	v.Memory[0x1A] = 0x80
	v.Memory[0x1B] = 0xF0
	v.Memory[0x1C] = 0x10
	v.Memory[0x1D] = 0xF0

	// 6
	v.Memory[0x1E] = 0xF0
	v.Memory[0x1F] = 0x80
	v.Memory[0x20] = 0xF0
	v.Memory[0x21] = 0x90
	v.Memory[0x22] = 0xF0

	// 7
	v.Memory[0x23] = 0xF0
	v.Memory[0x24] = 0x10
	v.Memory[0x25] = 0x20
	v.Memory[0x26] = 0x40
	v.Memory[0x27] = 0x40

	// 8
	v.Memory[0x28] = 0xF0
	v.Memory[0x29] = 0x90
	v.Memory[0x2A] = 0xF0
	v.Memory[0x2B] = 0x90
	v.Memory[0x2C] = 0xF0

	// 9
	v.Memory[0x2D] = 0xF0
	v.Memory[0x2E] = 0x90
	v.Memory[0x2F] = 0xF0
	v.Memory[0x30] = 0x10
	v.Memory[0x31] = 0xF0

	// A
	v.Memory[0x32] = 0xF0
	v.Memory[0x33] = 0x90
	v.Memory[0x34] = 0xF0
	v.Memory[0x35] = 0x90
	v.Memory[0x36] = 0x90

	// B
	v.Memory[0x37] = 0xE0
	v.Memory[0x38] = 0x90
	v.Memory[0x39] = 0xE0
	v.Memory[0x3A] = 0x90
	v.Memory[0x3B] = 0xE0

	// C
	v.Memory[0x3C] = 0xF0
	v.Memory[0x3D] = 0x80
	v.Memory[0x3E] = 0x80
	v.Memory[0x3F] = 0x80
	v.Memory[0x40] = 0xF0

	// D
	v.Memory[0x41] = 0xE0
	v.Memory[0x42] = 0x90
	v.Memory[0x43] = 0x90
	v.Memory[0x44] = 0x90
	v.Memory[0x45] = 0xE0

	// E
	v.Memory[0x46] = 0xF0
	v.Memory[0x47] = 0x80
	v.Memory[0x48] = 0xF0
	v.Memory[0x49] = 0x80
	v.Memory[0x4A] = 0xF0

	// F
	v.Memory[0x4B] = 0xF0
	v.Memory[0x4C] = 0x80
	v.Memory[0x4D] = 0xF0
	v.Memory[0x4E] = 0x80
	v.Memory[0x4F] = 0x80
}

func (v *Chip8) loadROM(rom []byte) {
	for i := 0; i < len(rom); i++ {
		v.Memory[0x200+i] = rom[i]
	}
}

// load file into a byte array
func (v *Chip8) LoadROMFromFile(path string) {
	rom, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	v.loadROM(rom)
}
