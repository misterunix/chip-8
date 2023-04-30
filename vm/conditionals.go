package vm

import "fmt"

func (v *Chip8) sne0x9000() {
	if v.Debug {
		v.DebugString += fmt.Sprintf("SNE V%X, V%X", v.x, v.y)

	}
	if v.Registers[v.x] != v.Registers[v.y] {
		v.PC += 2
	}
}

// Skip next instruction if Vx == NN
func (v *Chip8) se0x3000() {
	if v.Debug {
		v.DebugString += fmt.Sprintf("SE V%X, %02X", v.x, v.nn)

	}
	if v.Registers[v.x] == v.nn {
		v.PC += 2
	}
}

// Skip next instruction if Vx != NN
func (v *Chip8) sne0x4000() {
	if v.Debug {
		v.DebugString += fmt.Sprintf("SNE V%X, %02X", v.x, v.nn)

	}
	if v.Registers[v.x] != v.nn {
		v.PC += 2
	}
}

// Skip next instruction if Vx == Vy
func (v *Chip8) se0x5000() {
	if v.Debug {
		v.DebugString += fmt.Sprintf("SE V%X, V%X", v.x, v.y)

	}
	if v.Registers[v.x] == v.Registers[v.y] {
		v.PC += 2
	}
}

// Skip next instruction if key with the value of Vx is pressed
func (v *Chip8) skp0x009E() {
	if v.Debug {
		v.DebugString += fmt.Sprintf("SKP V%X", v.x)

	}
	if v.Keys[v.Registers[v.x]] == 1 {
		v.PC += 2
	}
}

// Skip next instruction if key with the value of Vx is not pressed
func (v *Chip8) sknp0x00A1() {
	if v.Debug {
		v.DebugString += fmt.Sprintf("SKNP V%X", v.x)

	}
	if v.Keys[v.Registers[v.x]] == 0 {
		v.PC += 2
	}
}
