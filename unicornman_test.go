package main

import "testing"

func BasicUnicornManager(t *testing.T) {
	um := UnicornManager{}
	um.init()
	for i := 0; i < 64; i++ {
		p := &Pixel{pos: uint(i), r: 255, b: 255, g: 255}
		um.SetPixelColor(p)
	}
	um.Stop()
}
