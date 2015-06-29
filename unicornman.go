package main

import (
	"github.com/DrItanium/unicornhat"
)

type UnicornManager struct {
}

type Pixel struct {
	pos uint
	r   byte
	b   byte
	g   byte
}

func (um *UnicornManager) init() {
	unicornhat.SetBrightness(unicornhat.DefaultBrightness())
	unicornhat.Init(64)
	unicornhat.ClearLEDBuffer()
}

func (um *UnicornManager) SetPixelColor(p Pixel) {
	unicornhat.SetPixelColor(p.pos, p.r, p.b, p.g)
	unicornhat.Show()
}

func (um *UnicornManager) Stop() {
	for i := 0; i < 64; i++ {
		unicornhat.SetPixelColor(uint(i), 0, 0, 0)
	}
	unicornhat.Show()
	unicornhat.Shutdown(0)
}
