package main

import "testing"
import "time"

func BasicUnicornManager(t *testing.T) {
	um := UnicornManager{}
	um.init()
	for i := 0; i < 64; i++ {
		p := &Pixel{pos: uint(i), r: 255, b: 255, g: 255}
		um.SetPixelColor(p)
		time.Sleep(1 * time.Second)
	}
	um.Stop()
}
