package main

import (
	"log"
	"time"
)

func main() {
	um := UnicornManager{}
	um.init()
	log.Println("Init Unicorn")
	for i := 0; i < 64; i++ {
		p := &Pixel{pos: uint(i), r: 255, b: 255, g: 255}
		log.Println("Pixel", p)
		um.SetPixelColor(p)
		time.Sleep(1 * time.Second)
	}
	um.Stop()
}
