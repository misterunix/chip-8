package main

import (
	"chip-8/vm"
	"flag"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 480
	scale        = 64
)

var v *vm.Chip8

type Game struct {
	keys []ebiten.Key
}

func NewGame() *Game {
	g := &Game{}

	return g
}

func (g *Game) Update() error {
	if v.ST > 0 {
		v.ST--
	}
	if v.DT > 0 {
		v.DT--
	}

	if ebiten.IsKeyPressed(ebiten.Key1) {
		v.Keys[0x01] = 1
	} else {
		v.Keys[0x01] = 0
	}
	if ebiten.IsKeyPressed(ebiten.Key2) {
		v.Keys[0x02] = 1
	} else {
		v.Keys[0x02] = 0
	}
	if ebiten.IsKeyPressed(ebiten.Key3) {
		v.Keys[0x03] = 1
	} else {
		v.Keys[0x03] = 0
	}
	if ebiten.IsKeyPressed(ebiten.Key4) {
		v.Keys[0x0C] = 1
	} else {
		v.Keys[0x0C] = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		v.Keys[0x04] = 1
	} else {
		v.Keys[0x04] = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		v.Keys[0x05] = 1
	} else {
		v.Keys[0x05] = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		v.Keys[0x06] = 1
	} else {
		v.Keys[0x06] = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		v.Keys[0x0D] = 1
	} else {
		v.Keys[0x0D] = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		v.Keys[0x07] = 1
	} else {
		v.Keys[0x07] = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		v.Keys[0x08] = 1
	} else {
		v.Keys[0x08] = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		v.Keys[0x09] = 1
	} else {
		v.Keys[0x09] = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyF) {
		v.Keys[0x0E] = 1
	} else {
		v.Keys[0x0E] = 0
	}

	if ebiten.IsKeyPressed(ebiten.KeyZ) {
		v.Keys[0x0A] = 1
	} else {
		v.Keys[0x0A] = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyX) {
		v.Keys[0x00] = 1
	} else {
		v.Keys[0x00] = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyC) {
		v.Keys[0x0B] = 1
	} else {
		v.Keys[0x0B] = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyV) {
		v.Keys[0x0F] = 1
	} else {
		v.Keys[0x0F] = 0
	}

	v.Execute()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//	if !vm.Draw {
	//		return
	//	}

	for y := 0; y < 32; y++ {
		for x := 0; x < 64; x++ {
			index := v.XYToIndex(uint8(x), uint8(y))
			t := v.Display[index]
			if t == 1 {
				vector.DrawFilledRect(screen, float32(x*10), float32(y*10), float32(10), float32(10), color.White, false)
			} else {
				vector.DrawFilledRect(screen, float32(x*10), float32(y*10), float32(10), float32(10), color.Black, false)
			}

		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 320
}

func main() {

	var romfile string

	flag.StringVar(&romfile, "rom", "", "ROM file to load")
	flag.Parse()

	if romfile == "" {
		log.Fatal("No ROM file specified")
	}

	v = vm.NewVM(0) // Create a new VM

	v.LoadROMFromFile(romfile) // Load the ROM file

	v.Reset() // Reset the VM, this will also clear the display

	go Run() // Run the VM in a goroutine so we can set the run speed.

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("CHIP-8")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}

func Run() {
	//
	for {
		v.Execute()

		//time.Sleep(1 * time.Second)
		//time.Sleep(500 * time.Millisecond)
		time.Sleep(1428 * time.Microsecond) // ~ 700 Hz
	}

}
