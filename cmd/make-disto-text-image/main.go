package main

import (
	"flag"
	"image/jpeg"
	"log"
	"os"
	"strings"

	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/inconsolata"

	textimgdisto "github.com/yinyin/go-textimgdisto"
)

func parseCommandParam() (textImageMaker textimgdisto.TextImageMaker, textContent string, distoCommands []distoMethod, outputFileName string, err error) {
	var fontName string
	var imageWidth, imageHeight int
	var fontSize float64
	flag.StringVar(&fontName, "fontName", "", "font name (one of: `mono` (default), `regular` or `inconsolata`)")
	flag.IntVar(&imageWidth, "width", 256, "image width")
	flag.IntVar(&imageHeight, "height", 80, "image height")
	flag.Float64Var(&fontSize, "fontSize", 72.0, "font size")
	flag.StringVar(&textContent, "text", "123789", "text content to draw")
	flag.StringVar(&outputFileName, "out", "disto-text-output.jpg", "output file name (should suffix with .jpg or .jpeg)")
	flag.Parse()
	if !strings.HasSuffix(outputFileName, ".jpg") && !strings.HasSuffix(outputFileName, ".jpeg") {
		log.Printf("WARN: output file name not suffixed with `.jpg` or `.jpeg`.")
	}
	switch strings.ToLower(fontName) {
	case "regular":
		var makerImpl *textimgdisto.OpenTypeTextImageMaker
		if makerImpl, err = textimgdisto.NewOpenTypeTextImageMakerWithFontData(
			goregular.TTF, imageWidth, imageHeight, fontSize); nil != err {
			return
		}
		textImageMaker = makerImpl
	case "inconsolata":
		makerImpl := textimgdisto.NewBasicTextImageMaker(inconsolata.Regular8x16, imageWidth, imageHeight)
		textImageMaker = makerImpl
	default:
		var makerImpl *textimgdisto.OpenTypeTextImageMaker
		if makerImpl, err = textimgdisto.NewOpenTypeTextImageMakerWithFontData(
			gomono.TTF, imageWidth, imageHeight, fontSize); nil != err {
			return
		}
		textImageMaker = makerImpl
	}
	args := flag.Args()
	for _, arg := range args {
		m, ok := parseDistoCommand(arg)
		if ok {
			distoCommands = append(distoCommands, m)
		} else {
			log.Printf("WARN: cannot parse into disto-command: %v", arg)
		}
	}
	if len(distoCommands) == 0 {
		log.Print("INFO: use default disto: cosh,0.16,6 cosv,0.07,6 blky,32,16 inv")
		distoCommands = append(distoCommands, &distoCosVShift{
			stepRadian: 0.16,
			ampValue:   6,
		}, &distoCosHShift{
			stepRadian: 0.07,
			ampValue:   6,
		}, &distoBlockyFlip7{
			blockWidth:  32,
			blockHeight: 16,
		}, &distoInvert{})
	}
	return
}

func main() {
	textImageMaker, textContent, distoCommands, outputFileName, err := parseCommandParam()
	if nil != err {
		log.Fatalf("cannot have command options: %v", err)
		return
	}
	dst, _, _, err := textImageMaker.NewTextImage(textContent)
	if nil != err {
		log.Fatalf("cannot make image: %v", err)
	}
	for _, distroM := range distoCommands {
		dst = distroM.disto(dst)
	}
	fp, err := os.Create(outputFileName)
	if nil != err {
		log.Fatalf("cannot open output file [%s]: %v", outputFileName, err)
		return
	}
	defer fp.Close()
	jpeg.Encode(fp, dst, &jpeg.Options{
		Quality: 50,
	})
}
