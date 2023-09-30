package gotextimgdistro

import (
	"image"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

type TextImageMaker interface {
	NewTextImage(t string) (dst *image.Gray, bearingX, bearingY fixed.Int26_6, err error)
}

type BasicTextImageMaker struct {
	fontFaceRef font.Face
	imageRect   image.Rectangle
	imageWidth  fixed.Int26_6
	imageHeight fixed.Int26_6
}

func NewBasicTextImageMaker(fontFaceRef font.Face, imageWidth, imageHeight int) (maker *BasicTextImageMaker) {
	maker = &BasicTextImageMaker{
		fontFaceRef: fontFaceRef,
		imageRect: image.Rectangle{
			Max: image.Point{
				X: imageWidth,
				Y: imageHeight,
			},
		},
		imageWidth:  fixed.I(imageWidth),
		imageHeight: fixed.I(imageHeight),
	}
	return
}

func (maker *BasicTextImageMaker) NewTextImage(t string) (dst *image.Gray, bearingX, bearingY fixed.Int26_6, err error) {
	face := maker.fontFaceRef
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

type OpenTypeTextImageMaker struct {
	fontRef     *opentype.Font
	fontSize    float64
	imageRect   image.Rectangle
	imageWidth  fixed.Int26_6
	imageHeight fixed.Int26_6
}

func NewOpenTypeTextImageMaker(fontRef *opentype.Font, imageWidth, imageHeight int, fontSize float64) (maker *OpenTypeTextImageMaker) {
	maker = &OpenTypeTextImageMaker{
		fontRef:  fontRef,
		fontSize: fontSize,
		imageRect: image.Rectangle{
			Max: image.Point{
				X: imageWidth,
				Y: imageHeight,
			},
		},
		imageWidth:  fixed.I(imageWidth),
		imageHeight: fixed.I(imageHeight),
	}
	return
}

func NewOpenTypeTextImageMakerWithFontData(fontData []byte, imageWidth, imageHeight int, fontSize float64) (maker *OpenTypeTextImageMaker, err error) {
	fontRef, err := opentype.Parse(fontData)
	if nil != err {
		return
	}
	maker = NewOpenTypeTextImageMaker(fontRef, imageWidth, imageHeight, fontSize)
	return
}

func (maker *OpenTypeTextImageMaker) NewTextImage(t string) (dst *image.Gray, bearingX, bearingY fixed.Int26_6, err error) {
	face, err := opentype.NewFace(maker.fontRef, &opentype.FaceOptions{
		Size:    maker.fontSize,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if nil != err {
		return
	}
	defer face.Close()
	basicMaker := BasicTextImageMaker{
		fontFaceRef: face,
		imageRect:   maker.imageRect,
		imageWidth:  maker.imageWidth,
		imageHeight: maker.imageHeight,
	}
	return basicMaker.NewTextImage(t)
}
