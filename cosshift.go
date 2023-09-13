package gotextimgdistro

import (
	"image"
	"math"
)

func CosHShift(srcImg *image.Gray, stepRadian, ampValue float64) (dstImg *image.Gray) {
	width := srcImg.Rect.Max.X
	height := srcImg.Rect.Max.Y
	dstImg = image.NewGray(srcImg.Rect)
	radiusValue := stepRadian
	for y := 0; y < height; y++ {
		radiusValue += stepRadian
		if radiusValue > (math.Pi * 2) {
			radiusValue = 0
		}
		xShift := int(math.Cos(radiusValue) * ampValue)
		for x := 0; x < width; x++ {
			if shftX := x + xShift; (shftX >= 0) && (shftX < width) {
				dstImg.SetGray(shftX, y, srcImg.GrayAt(x, y))
			}
		}
	}
	return
}

func CosVShift(srcImg *image.Gray, stepRadian, ampValue float64) (dstImg *image.Gray) {
	width := srcImg.Rect.Max.X
	height := srcImg.Rect.Max.Y
	dstImg = image.NewGray(srcImg.Rect)
	radiusValue := stepRadian
	for x := 0; x < width; x++ {
		radiusValue += stepRadian
		if radiusValue > (math.Pi * 2) {
			radiusValue = 0
		}
		yShift := int(math.Cos(radiusValue) * ampValue)
		for y := 0; y < height; y++ {
			if shftY := y + yShift; (shftY >= 0) && (shftY < height) {
				dstImg.SetGray(x, shftY, srcImg.GrayAt(x, y))
			}
		}
	}
	return
}
