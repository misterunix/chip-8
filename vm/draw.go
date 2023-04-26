package vm

import "fmt"

// Draw sprite at Vx, Vy with width N
func (v *Chip8) drw0xD000() {
	if v.debug {
		v.DebugString += fmt.Sprintf("DRW\tV%X, V%X, %X\n", v.x, v.y, v.n)

	}
	v.Draw = true
	yi := v.Registers[v.y]
	xi := v.Registers[v.x]
	v.Registers[0xF] = 0

	for k := 0; k < int(v.n); k++ {
		iv := v.I + uint16(k)
		if iv >= 4096 {
			fmt.Println("ERROR: v.I + k out of bounds", iv)
			v.Reset()
		}
		q := v.Memory[iv]
		for j := 0; j < 8; j++ {
			b := (q >> (7 - j)) & 0x1
			if b == 1 {
				tindex := v.XYToIndex((uint8(int(xi)+j) % 64), (uint8(int(yi)+k) % 32))
				if tindex >= 2048 {
					fmt.Println("ERROR: tindex out of bounds", tindex, (uint8(int(xi)+j) % 64), (uint8(int(yi)+k) % 32))
					v.Reset()
				}
				v.Display[tindex] ^= 1
				if v.Display[tindex] == 0 {
					v.Registers[0xF] = 1
				}

			}
		}
	}

	/*
		for i := yi; i < yi+v.n; i++ {
			for j := xi; j < xi+8; j++ {
				bit := (v.Memory[v.I+uint16(i)] >> (7 - (j - xi))) & 0x1
				if bit == 1 {
					index := v.XYToIndex(uint8(j), uint8(i))
					if v.Display[index] == 1 {
						v.Registers[0xF] = 1
					}
					v.Display[index] ^= 1
					if v.Display[index] == 1 {
						fmt.Print("X")
					} else {
						fmt.Print(" ")
					}

				}
			}
			fmt.Println()
		}

		fmt.Println()
		fmt.Println()
	*/
}
