package vm

import "fmt"

// Set Vx = Vy
func (v *Chip8) ld8000() {
	if v.Debug {
		v.DebugString += fmt.Sprintf("LD V%X(%x), V%X(%x)", v.x, v.FetchRegister(v.x), v.y, v.FetchRegister(v.y))
	}
	tv := v.FetchRegister(v.y)
	v.StoreRegister(v.x, tv)
	//v.Registers[v.x] = v.Registers[v.y]
}

// Set Vx = Vx OR Vy
func (v *Chip8) or8001() {
	if v.Debug {
		v.DebugString += fmt.Sprintf("OR V%X(%x), V%X(%x)", v.x, v.FetchRegister(v.x), v.y, v.FetchRegister(v.y))
	}
	tvy := v.FetchRegister(v.y)
	tvx := v.FetchRegister(v.x)
	tv := tvx | tvy
	v.StoreRegister(v.x, tv)
	//v.Registers[v.x] |= v.Registers[v.y]
}

// Set Vx = Vx AND Vy
func (v *Chip8) and8002() {
	if v.Debug {
		v.DebugString += fmt.Sprintf("AND V%X(%x), V%X(%x)", v.x, v.FetchRegister(v.x), v.y, v.FetchRegister(v.y))
	}
	tvy := v.FetchRegister(v.y)
	tvx := v.FetchRegister(v.x)
	tv := tvx & tvy
	v.StoreRegister(v.x, tv)
	//v.Registers[v.x] &= v.Registers[v.y]
}

// Set Vx = Vx XOR Vy
func (v *Chip8) xor8003() {
	if v.Debug {
		v.DebugString += fmt.Sprintf("XOR V%X(%x), V%X(%x)", v.x, v.FetchRegister(v.x),
			v.y, v.FetchRegister(v.y))
	}
	tvy := v.FetchRegister(v.y)
	tvx := v.FetchRegister(v.x)
	tv := tvx ^ tvy
	v.StoreRegister(v.x, tv)
	//v.Registers[v.x] ^= v.Registers[v.y]
}

// Set Vx = Vx + Vy, set VF = carry
func (v *Chip8) add8004() {
	if v.Debug {
		v.DebugString += fmt.Sprintf("ADD V%X(%x), V%X(%x)", v.x, v.FetchRegister(v.x),
			v.y, v.FetchRegister(v.y))
	}
	tx := v.FetchRegister(v.x)
	ty := v.FetchRegister(v.y)
	if int(tx)+int(ty) > 255 {
		v.StoreRegister(0xF, 1)
		//v.Registers[0xF] = 1
	} else {
		v.StoreRegister(0xF, 0)
		//v.Registers[0xF] = 0
	}
	v.StoreRegister(v.x, tx+ty)
	//v.Registers[v.x] += v.Registers[v.y]
}

// Set Vx = Vx - Vy, set VF = NOT borrow
func (v *Chip8) sub8005() {
	if v.Debug {
		v.DebugString += fmt.Sprintf("SUB V%X(%x), V%X(%x)", v.x, v.FetchRegister(v.x),
			v.y, v.FetchRegister(v.y))
	}
	tx := v.FetchRegister(v.x)
	ty := v.FetchRegister(v.y)
	if int(tx) > int(ty) {
		//v.StoreRegister(0xF, 1)
		v.Registers[0xF] = 1
	} else {
		v.StoreRegister(0xF, 0)
		//v.Registers[0xF] = 0
	}
	v.StoreRegister(v.x, tx-ty)
	//v.Registers[v.x] -= v.Registers[v.y]
}

// Set Vx = Vx SHR 1
func (v *Chip8) shr8006() {
	if v.Debug {
		v.DebugString += fmt.Sprintf("SHR V%X(%x) {,V%X(%x)}", v.x, v.FetchRegister(v.x), v.x, v.FetchRegister(v.x))
	}
	tx := v.FetchRegister(v.x) & 0x1
	v.StoreRegister(0xF, tx)
	v.StoreRegister(v.x, tx/2)
	//v.Registers[v.x] = v.Registers[v.y]
	//v.Registers[0xF] = v.Registers[v.x] & 0x1
	//v.Registers[v.x] = v.Registers[v.x] / 2

	//v.Registers[0xF] = v.Registers[v.x] & 0x1
	//v.Registers[v.x] /= 2
}

// Set Vx = Vy - Vx, set VF = NOT borrow
func (v *Chip8) subn8007() {
	if v.Debug {
		v.DebugString += fmt.Sprintf("SUBN V%X(%x), V%X(%x)", v.x, v.FetchRegister(v.x),
			v.y, v.FetchRegister(v.y))
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
	if v.Debug {
		v.DebugString += fmt.Sprintf("SHL V%X(%X) {, V%X(%X)}", v.x, v.FetchRegister(v.x), v.y, v.FetchRegister(v.y))
	}
	//v.Registers[v.x] = v.Registers[v.y]
	tx := v.FetchRegister(v.x)
	if tx&0x80 == 128 {
		v.StoreRegister(0xF, 1)
		//v.Registers[0xF] = 1
	} else {
		v.StoreRegister(0xF, 0)
		//v.Registers[0xF] = 0
	}
	//v.Registers[0xF] = v.Registers[v.x] & 0x80 >> 7
	v.StoreRegister(v.x, tx*2)
	//v.Registers[v.x] = v.Registers[v.x] * 2
}

// Set I = I + Vx
func (v *Chip8) addi0x001E() {
	if v.Debug {
		v.DebugString += fmt.Sprintf("ADD I(%04x), V%X(%x)", v.I, v.x, v.FetchRegister(v.x))
	}
	tx := v.FetchRegister(v.x)
	v.I += uint16(tx)
	if v.I > 0xFFF {
		if v.Debug {
			v.DebugString += fmt.Sprintf("I > 0xFFF, I = %X", v.I)
		}
		v.Reset()
	}
}
