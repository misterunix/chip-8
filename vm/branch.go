package vm

import "fmt"

// Return from subroutine
func (v *Chip8) ret0x00EE() {
	if v.debug {
		v.DebugString += "RET"
	}
	v.SP--
	v.PC = v.Stack[v.SP]
}

// Jump to location nnn
func (v *Chip8) jp0x1000() {
	if v.debug {
		v.DebugString += fmt.Sprintf("JP %04X", v.nnn)
	}
	v.PC = v.nnn
}

// Call subroutine
func (v *Chip8) call0x2000() {
	if v.debug {
		v.DebugString += fmt.Sprintf("CALL %04X", v.nnn)
	}
	v.Stack[v.SP] = v.PC
	v.SP++
	v.PC = v.nnn
}

// Jump to location nnn + V0
func (v *Chip8) jp0xB000() {
	if v.debug {
		v.DebugString += fmt.Sprintf("JP V0, %04X", v.nnn)
	}
	v.PC = v.nnn + uint16(v.Registers[0])
}
