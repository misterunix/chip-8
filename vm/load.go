package vm

import (
	"fmt"
)

// Load I with address nnn
func (v *Chip8) ldi0xA000() {
	if v.debug {
		v.DebugString += fmt.Sprintf("LD I, %04X", v.nnn)

	}
	v.I = v.nnn
}

// Load Vx with byte nn
func (v *Chip8) ld0x6000() {
	if v.debug {
		v.DebugString += fmt.Sprintf("LD V%X, %02X", v.x, v.nn)

	}
	v.Registers[v.x] = v.nn
}

// Set Vx = random byte AND kk
func (v *Chip8) ldrnd0xC000() {
	if v.debug {
		v.DebugString += fmt.Sprintf("RND V%X, %02X", v.x, v.nn)

	}
	v.Registers[v.x] = uint8(v.rnd.Intn(256)) & v.nn
	//fmt.Println(v.Registers[v.x])
}

// Set Vx = delay timer value
func (v *Chip8) ldvtimer0x0007() {
	if v.debug {
		v.DebugString += fmt.Sprintf("LD V%X, DT", v.x)

	}
	v.Registers[v.x] = v.DT
}

// Set delay timer = Vx
func (v *Chip8) lddtvx0x0015() {
	if v.debug {
		v.DebugString += fmt.Sprintf("LD DT, V%X", v.x)

	}
	v.DT = v.Registers[v.x]
}

// Set sound timer = Vx
func (v *Chip8) ldst0x0018() {
	if v.debug {
		v.DebugString += fmt.Sprintf("LD ST, V%X", v.x)

	}
	v.ST = v.Registers[v.x]
}

// Set I = location of sprite for digit Vx
func (v *Chip8) ldi0x0029() {
	if v.debug {
		v.DebugString += fmt.Sprintf("LD F, V%X", v.x)

	}
	v.I = uint16(v.Registers[v.x]) * 5
}

// Store BCD representation of Vx in memory locations I, I+1, and I+2
func (v *Chip8) bcd0x0033() {
	if v.debug {
		v.DebugString += fmt.Sprintf("LD B. V%X", v.x)

	}
	v.Memory[v.I] = v.Registers[v.x] / 100
	v.Memory[v.I+1] = (v.Registers[v.x] / 10) % 10
	v.Memory[v.I+2] = (v.Registers[v.x] % 100) % 10
}

// Store registers V0 through Vx in memory starting at location I
func (v *Chip8) lsmemv0x0055() {
	if v.debug {
		v.DebugString += fmt.Sprintf("LD [I], V%X", v.x)

	}
	for i := 0; i <= int(v.x); i++ {
		v.Memory[v.I+uint16(i)] = v.Registers[i]
	}
}

// Read registers V0 through Vx from memory starting at location I
func (v *Chip8) ldvmem0x0065() {
	if v.debug {
		v.DebugString += fmt.Sprintf("LD V%X, [I]", v.x)

	}
	for i := 0; i <= int(v.x); i++ {
		v.Registers[i] = v.Memory[v.I+uint16(i)]
	}
}
