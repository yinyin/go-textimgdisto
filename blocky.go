package gotextimgdistro

import (
	"image"
	"image/color"
)

func BlockyFlip7(srcImg *image.Gray, blockWidth, blockHeight int) (dstImg *image.Gray) {
	width := srcImg.Rect.Max.X
	height := srcImg.Rect.Max.Y
	dstImg = image.NewGray(srcImg.Rect)
	blockYStrip := (width + 1) / blockWidth
	hCountdown := blockHeight
	blockYOffset := 0
	flipTarget := 3
	for y := 0; y < height; y++ {
		hCountdown--
		if hCountdown == 0 {
			hCountdown = blockHeight
			blockYOffset += blockYStrip
			flipTarget = (flipTarget + 3) % 7
		}
		blockIndex := blockYOffset
		wCountdown := blockWidth
		doFlip := (blockIndex % 7) == flipTarget
		for x := 0; x < width; x++ {
			wCountdown--
			if wCountdown == 0 {
				blockIndex++
				wCountdown = blockWidth
				doFlip = (blockIndex % 7) == flipTarget
			}
			if doFlip {
				c := color.Gray{
					Y: 255 - srcImg.GrayAt(x, y).Y,
				}
				dstImg.SetGray(x, y, c)
			} else {
				dstImg.SetGray(x, y, srcImg.GrayAt(x, y))
			}
		}
	}
	return
}
