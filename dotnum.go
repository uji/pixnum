package pixnum

import (
	"image"
	"image/color"
)

func NewImage(number int, color color.Color) image.Image {
	if number > 9 || number < 0 {
		panic("number is not one digit")
	}

	rgba := image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{3, 5},
	})

	ps := getPoints(number)

	for _, p := range ps {
		rgba.Set(p.x, p.y, color)
	}

	return rgba
}
