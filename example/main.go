package main

import (
	"fmt"
	"image/color"
	"image/png"
	"os"

	"github.com/uji/pixnum"
)

func main() {
	for i := 0; i < 10; i++ {
		createImage(i)
	}
}

func createImage(n int) {
	img := pixnum.NewImage(n, color.Black)

	path := fmt.Sprintf("image%d.png", n)

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		panic(err)
	}
}
