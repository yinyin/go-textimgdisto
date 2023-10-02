package gotextimgdistro

import (
	"image"
	"math"
)

const (
	tanShiftMinRadian = -math.Pi / 4
	tanShiftMaxRadian = math.Pi / 4
)

func tanShiftMakeInitRadian(initRadian float64) float64 {
	if initRadian < tanShiftMinRadian {
		return tanShiftMinRadian
	} else if initRadian > tanShiftMaxRadian {
		return tanShiftMaxRadian
	}
	return initRadian
}

func TanHShift(srcImg *image.Gray, initRadian, stepRadian, ampValue float64) (dstImg *image.Gray) {
	width := srcImg.Rect.Max.X
	height := srcImg.Rect.Max.Y
	dstImg = image.NewGray(srcImg.Rect)
	radiusValue := tanShiftMakeInitRadian(initRadian)
	flipValue := 1.0
	for y := 0; y < height; y++ {
		radiusValue += stepRadian
		if radiusValue > (math.Pi / 4) {
			radiusValue = -math.Pi / 4
			flipValue = -flipValue
		}
		currTan := math.Tan(radiusValue) * flipValue
		xShift := int(currTan * ampValue)
		for x := 0; x < width; x++ {
			if shftX := x + xShift; (shftX >= 0) && (shftX < width) {
				dstImg.SetGray(shftX, y, srcImg.GrayAt(x, y))
			}
		}
	}
	return
}

func TanVShift(srcImg *image.Gray, initRadian, stepRadian, ampValue float64) (dstImg *image.Gray) {
	width := srcImg.Rect.Max.X
	height := srcImg.Rect.Max.Y
	dstImg = image.NewGray(srcImg.Rect)
	radiusValue := tanShiftMakeInitRadian(initRadian)
	flipValue := 1.0
	for x := 0; x < width; x++ {
		radiusValue += stepRadian
		if radiusValue > (math.Pi / 4) {
			radiusValue = -math.Pi / 4
			flipValue = -flipValue
		}
		currTan := math.Tan(radiusValue) * flipValue
		yShift := int(currTan * ampValue)
		for y := 0; y < height; y++ {
			if shftY := y + yShift; (shftY >= 0) && (shftY < height) {
				dstImg.SetGray(x, shftY, srcImg.GrayAt(x, y))
			}
		}
	}
	return
}
