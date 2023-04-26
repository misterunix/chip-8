package vm

func (v *Chip8) IndexToXY(i int) (uint8, uint8) {
	var x, y int
	if i >= v.DisplaySize {
		v.Reset()
	}
	x = i % int(v.ScreenSize[0])
	y = i / int(v.ScreenSize[0])
	return uint8(x), uint8(y)
}

func (v *Chip8) XYToIndex(x, y uint8) int {
	i := int(v.ScreenSize[0])*int(y) + int(x)
	if i >= v.DisplaySize {
		v.Reset()
	}
	return i
}
