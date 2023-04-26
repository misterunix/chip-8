package vm

import "fmt"

func (v *Chip8) Execute() {
	v.Draw = false
	v.OpCode = uint16(v.Memory[v.PC])<<8 | uint16(v.Memory[v.PC+1]) // Get the opcode
	if v.debug {
		v.DebugString = fmt.Sprintf("%04X %04X ", v.PC, v.OpCode)
	}
	v.PC += 2
	// x register is the first nibble of the opcode
	v.x = uint8((v.OpCode & 0x0F00) >> 8) // Decode Vx register
	// y register is the second nibble of the opcode
	v.y = uint8((v.OpCode & 0x00F0) >> 4) // Decode Vy register
	//  n is the last nibble of the opcode
	v.n = uint8(v.OpCode & 0x000F)
	// nn is the last two bytes of the opcode
	v.nn = uint8(v.OpCode & 0x00FF) // Decode NN
	// nnn is the last three bytes of the opcode
	v.nnn = uint16(v.OpCode & 0x0FFF) // Decode NNN
	if v.nnn > 0x0FFF {
		if v.debug {
			v.DebugString += fmt.Sprintf("Invalid opcode: %04X", v.OpCode)
		}
		v.Reset()
	}
	switch v.OpCode & 0xF000 {

	case 0x0000:
		switch v.OpCode & 0x00FF {
		case 0x00E0: // Clear screen
			if v.debug {
				v.DebugString += "CLS"
			}
			v.clearScreen()
		case 0x00EE: // Return from subroutine
			v.ret0x00EE()
		}

	case 0x1000:
		v.jp0x1000()
	case 0x2000: // Call subroutine
		v.call0x2000()
	case 0x3000: // Skip next instruction if Vx == NN
		v.se0x3000()
	case 0x4000: // Skip next instruction if Vx != NN
		v.sne0x4000()
	case 0x5000: // Skip next instruction if Vx == Vy
		v.se0x5000()
	case 0x6000: // Load Vx with byte nn
		v.ld0x6000()
	case 0x7000:
		if v.debug {
			v.DebugString += fmt.Sprintf("ADD V%X, %02X", v.x, v.nn)
		}
		v.Registers[v.x] += v.nn
	case 0x8000: // Arithmetic
		switch v.OpCode & 0x000F {
		case 0x0000: // Set Vx = Vy
			v.ld8000()
		case 0x0001: // Set Vx = Vx OR Vy
			v.or8001()
		case 0x0002: // Set Vx = Vx AND Vy
			v.and8002()
		case 0x0003: // Set Vx = Vx XOR Vy
			v.xor8003()
		case 0x0004: // Set Vx = Vx + Vy, set VF = carry
			v.add8004()
		case 0x0005: // Set Vx = Vx - Vy, set VF = NOT borrow
			v.sub8005()
		case 0x0006: // Set Vx = Vx SHR 1
			v.shr8006()
		case 0x0007: // Set Vx = Vy - Vx, set VF = NOT borrow
			v.subn8007()
		case 0x000E: // Set Vx = Vx SHL 1
			v.shl800E()
		}
	case 0x9000: // Skip next instruction if Vx != Vy
		v.sne0x9000()
	case 0xA000:
		v.ldi0xA000()
	case 0xB000: // Jump to location nnn + V0
		v.jp0xB000()
	case 0xC000: // Set Vx = random byte AND kk
		v.ldrnd0xC000()
	case 0xD000: // // Draw sprite at Vx, Vy with width N
		v.drw0xD000()
	case 0xE000:
		switch v.nn {
		case 0x009E: // Skip next instruction if key with the value of Vx is pressed
			v.skp0x009E()
		case 0x00A1: // Skip next instruction if key with the value of Vx is not pressed
			v.sknp0x00A1()
		}
	case 0xF000:
		switch v.nn {
		case 0x0007: // Set Vx = delay timer value
			v.ldvtimer0x0007()
		case 0x0015: // Set delay timer = Vx
			v.lddtvx0x0015()
		case 0x0018: // Set sound timer = Vx
			v.ldst0x0018()
		case 0x001E: // Set I = I + Vx
			v.addi0x001E()
		case 0x0029: // Set I = location of sprite for digit Vx
			v.ldi0x0029()
		case 0x0033: // Store BCD representation of Vx in memory locations I, I+1, and I+2
			v.bcd0x0033()
		case 0x0055: // Store registers V0 through Vx in memory starting at location I
			v.lsmemv0x0055()
		case 0x0065: // Read registers V0 through Vx from memory starting at location I
			v.ldvmem0x0065()
		}
	}

}
