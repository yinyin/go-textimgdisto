package main

import (
	"image"
	"log"
	"strconv"
	"strings"

	textimgdisto "github.com/yinyin/go-textimgdisto"
)

type distoMethod interface {
	disto(dst *image.Gray) *image.Gray
}

type distoCosHShift struct {
	initRadian float64
	stepRadian float64
	ampValue   float64
}

func (m *distoCosHShift) disto(dst *image.Gray) *image.Gray {
	return textimgdisto.CosHShift(dst, m.initRadian, m.stepRadian, m.ampValue)
}

type distoCosVShift struct {
	initRadian float64
	stepRadian float64
	ampValue   float64
}

func (m *distoCosVShift) disto(dst *image.Gray) *image.Gray {
	return textimgdisto.CosVShift(dst, m.initRadian, m.stepRadian, m.ampValue)
}

type distoTanHShift struct {
	initRadian float64
	stepRadian float64
	ampValue   float64
}

func (m *distoTanHShift) disto(dst *image.Gray) *image.Gray {
	return textimgdisto.TanHShift(dst, m.initRadian, m.stepRadian, m.ampValue)
}

type distoTanVShift struct {
	initRadian float64
	stepRadian float64
	ampValue   float64
}

func (m *distoTanVShift) disto(dst *image.Gray) *image.Gray {
	return textimgdisto.TanVShift(dst, m.initRadian, m.stepRadian, m.ampValue)
}

type distoBlockyFlip7 struct {
	blockWidth  int
	blockHeight int
}

func (m *distoBlockyFlip7) disto(dst *image.Gray) *image.Gray {
	return textimgdisto.BlockyFlip7(dst, m.blockWidth, m.blockHeight)
}

type distoInvert struct {
}

func (m *distoInvert) disto(dst *image.Gray) *image.Gray {
	return textimgdisto.Invert(dst)
}

func parseDistoCommandTrigoShift(aux []string, auxCnt int, cmdText string) (initRadian, stepRadian, ampValue float64, ok bool) {
	if auxCnt != 4 {
		log.Printf("WARN: expect `%s,0.0,0.12,6.0`; got: %v", cmdText, aux)
		return
	}
	var err error
	if initRadian, err = strconv.ParseFloat(aux[1], 64); nil != err {
		log.Printf("WARN: cannot parse %s init-radian [%v]: %v", cmdText, aux[1], err)
		return
	}
	if stepRadian, err = strconv.ParseFloat(aux[2], 64); nil != err {
		log.Printf("WARN: cannot parse %s step-radian [%v]: %v", cmdText, aux[2], err)
		return
	}
	if ampValue, err = strconv.ParseFloat(aux[3], 64); nil != err {
		log.Printf("WARN: cannot parse %s amp-value [%v]: %v", cmdText, aux[3], err)
		return
	}
	ok = true
	return
}

func parseDistoCommand(cmd string) (m distoMethod, ok bool) {
	aux := strings.Split(cmd, ",")
	auxCnt := len(aux)
	if auxCnt < 1 {
		return
	}
	switch aux[0] {
	case "cosh":
		initRadian, stepRadian, ampValue, isParsed := parseDistoCommandTrigoShift(aux, auxCnt, "cosh")
		if !isParsed {
			return
		}
		m = &distoCosHShift{
			initRadian: initRadian,
			stepRadian: stepRadian,
			ampValue:   ampValue,
		}
		ok = true
	case "cosv":
		initRadian, stepRadian, ampValue, isParsed := parseDistoCommandTrigoShift(aux, auxCnt, "cosv")
		if !isParsed {
			return
		}
		m = &distoCosVShift{
			initRadian: initRadian,
			stepRadian: stepRadian,
			ampValue:   ampValue,
		}
		ok = true
	case "tanh":
		initRadian, stepRadian, ampValue, isParsed := parseDistoCommandTrigoShift(aux, auxCnt, "tanh")
		if !isParsed {
			return
		}
		m = &distoTanHShift{
			initRadian: initRadian,
			stepRadian: stepRadian,
			ampValue:   ampValue,
		}
		ok = true
	case "tanv":
		initRadian, stepRadian, ampValue, isParsed := parseDistoCommandTrigoShift(aux, auxCnt, "tanv")
		if !isParsed {
			return
		}
		m = &distoTanVShift{
			initRadian: initRadian,
			stepRadian: stepRadian,
			ampValue:   ampValue,
		}
		ok = true
	case "blky":
		if auxCnt != 3 {
			log.Print("WARN: expect `blky,32,16`; got:", aux)
			return
		}
		blockWidth, err := strconv.ParseInt(aux[1], 10, 32)
		if nil != err {
			log.Printf("WARN: cannot parse blky block-width [%v]: %v", aux[1], err)
			return
		}
		blockHeight, err := strconv.ParseInt(aux[2], 10, 32)
		if nil != err {
			log.Printf("WARN: cannot parse blky block-height [%v]: %v", aux[2], err)
			return
		}
		m = &distoBlockyFlip7{
			blockWidth:  int(blockWidth),
			blockHeight: int(blockHeight),
		}
		ok = true
	case "inv":
		m = &distoInvert{}
		ok = true
	}
	return
}
