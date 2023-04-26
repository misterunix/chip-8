package vm

import "fmt"

// Set Vx = Vy
func (v *Chip8) ld8000() {
	if v.debug {
		v.DebugString += fmt.Sprintf("LD V%X, V%X", v.x, v.y)
	}
	v.Registers[v.x] = v.Registers[v.y]
}

// Set Vx = Vx OR Vy
func (v *Chip8) or8001() {
	if v.debug {
		v.DebugString += fmt.Sprintf("OR V%X, V%X", v.x, v.y)
	}
	v.Registers[v.x] |= v.Registers[v.y]
}

// Set Vx = Vx AND Vy
func (v *Chip8) and8002() {
	if v.debug {
		v.DebugString += fmt.Sprintf("AND V%X, V%X", v.x, v.y)
	}
	v.Registers[v.x] &= v.Registers[v.y]
}

// Set Vx = Vx XOR Vy
func (v *Chip8) xor8003() {
	if v.debug {
		v.DebugString += fmt.Sprintf("XOR V%X, V%X", v.x, v.y)
	}
	v.Registers[v.x] ^= v.Registers[v.y]
}

// Set Vx = Vx + Vy, set VF = carry
func (v *Chip8) add8004() {
	if v.debug {
		v.DebugString += fmt.Sprintf("ADD V%X, V%X", v.x, v.y)
	}
	if int(v.Registers[v.x])+int(v.Registers[v.y]) > 255 {
		v.Registers[0xF] = 1
	} else {
		v.Registers[0xF] = 0
	}
	v.Registers[v.x] += v.Registers[v.y]
}

// Set Vx = Vx - Vy, set VF = NOT borrow
func (v *Chip8) sub8005() {
	if v.debug {
		v.DebugString += fmt.Sprintf("SUB V%X, V%X", v.x, v.y)
	}
	if int(v.Registers[v.x]) > int(v.Registers[v.y]) {
		v.Registers[0xF] = 1
	} else {
		v.Registers[0xF] = 0
	}
	v.Registers[v.x] -= v.Registers[v.y]
}

// Set Vx = Vx SHR 1
func (v *Chip8) shr8006() {
	if v.debug {
		v.DebugString += fmt.Sprintf("SHR V%X {,V%X}", v.x, v.x)
	}
	//v.Registers[v.x] = v.Registers[v.y]
	v.Registers[0xF] = v.Registers[v.x] & 0x1
	v.Registers[v.x] = v.Registers[v.x] / 2

	//v.Registers[0xF] = v.Registers[v.x] & 0x1
	//v.Registers[v.x] /= 2
}

// Set Vx = Vy - Vx, set VF = NOT borrow
func (v *Chip8) subn8007() {
	if v.debug {
		v.DebugString += fmt.Sprintf("SUBN V%X, V%X", v.x, v.y)
	}
	if int(v.Registers[v.y]) > int(v.Registers[v.x]) {
		v.Registers[0xF] = 1
	} else {
		v.Registers[0xF] = 0
	}
	v.Registers[v.x] = v.Registers[v.y] - v.Registers[v.x]
}

// Set Vx = Vx SHL 1
func (v *Chip8) shl800E() {
	if v.debug {
		v.DebugString += fmt.Sprintf("SHL V%X(%X) {, V%X(%X)}", v.x, v.Registers[v.x], v.y, v.Registers[v.y])
	}
	//v.Registers[v.x] = v.Registers[v.y]
	if v.Registers[v.x]&0x80 == 128 {
		v.Registers[0xF] = 1
	} else {
		v.Registers[0xF] = 0
	}
	//v.Registers[0xF] = v.Registers[v.x] & 0x80 >> 7
	v.Registers[v.x] = v.Registers[v.x] * 2

	// v.Registers[0xF] = v.Registers[v.x] >> 7
	// v.Registers[v.x] *= 2
}

// Set I = I + Vx
func (v *Chip8) addi0x001E() {
	if v.debug {
		v.DebugString += fmt.Sprintf("ADD I, V%X", v.x)
	}
	v.I += uint16(v.Registers[v.x])
	if v.I > 0xFFF {
		if v.debug {
			v.DebugString += fmt.Sprintf("I > 0xFFF, I = %X", v.I)
		}
		v.Reset()
	}
}
