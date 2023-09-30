package gotextimgdistro

import (
	"image"
	"image/color"
)

func Invert(srcImg *image.Gray) (dstImg *image.Gray) {
	width := srcImg.Rect.Max.X
	height := srcImg.Rect.Max.Y
	dstImg = image.NewGray(srcImg.Rect)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dstImg.SetGray(x, y, color.Gray{
				Y: 255 - srcImg.GrayAt(x, y).Y,
			})
		}
	}
	return
}
