package gotextimgdistro

import (
	"image"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
)

type TextImageMaker struct {
	fontRef     *sfnt.Font
	imageRect   image.Rectangle
	imageWidth  fixed.Int26_6
	imageHeight fixed.Int26_6
	fontSize    float64
}

func NewTextImageMaker(fontRef *sfnt.Font, imageWidth, imageHeight int, fontSize float64) (maker *TextImageMaker) {
	maker = &TextImageMaker{
		fontRef: fontRef,
		imageRect: image.Rectangle{
			Max: image.Point{
				X: imageWidth,
				Y: imageHeight,
			},
		},
		imageWidth:  fixed.I(imageWidth),
		imageHeight: fixed.I(imageHeight),
		fontSize:    fontSize,
	}
	return
}

func (maker *TextImageMaker) NewTextImage(t string) (dst *image.Gray, bearingX, bearingY fixed.Int26_6, err error) {
	face, err := opentype.NewFace(maker.fontRef, &opentype.FaceOptions{
		Size:    maker.fontSize,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if nil != err {
		return
	}
	defer face.Close()
	dst = image.NewGray(maker.imageRect)
	d := font.Drawer{
		Dst:  dst,
		Src:  image.White,
		Face: face,
	}
	textAdvance := d.MeasureString(t)
	var maxAscent, maxDescent fixed.Int26_6
	for _, ch := range t {
		bnd, _, _ := face.GlyphBounds(ch)
		if bnd.Min.Y < maxAscent {
			maxAscent = bnd.Min.Y
		}
		if bnd.Max.Y > maxDescent {
			maxDescent = bnd.Max.Y
		}
	}
	textHeight := maxDescent - maxAscent
	bearingX = (maker.imageWidth - textAdvance) / 2
	bearingY = (maker.imageHeight - textHeight) / 2
	d.Dot = fixed.Point26_6{
		X: bearingX,
		Y: bearingY - maxAscent,
	}
	d.DrawString(t)
	return
}
