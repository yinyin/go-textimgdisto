package main

import (
	"image"
	"log"
	"strconv"
	"strings"

	textimgdisto "github.com/yinyin/go-textimgdisto"
)

type distroMethod interface {
	distro(dst *image.Gray) *image.Gray
}

type distroCosHShift struct {
	stepRadian float64
	ampValue   float64
}

func (m *distroCosHShift) distro(dst *image.Gray) *image.Gray {
	return textimgdisto.CosHShift(dst, m.stepRadian, m.ampValue)
}

type distroCosVShift struct {
	stepRadian float64
	ampValue   float64
}

func (m *distroCosVShift) distro(dst *image.Gray) *image.Gray {
	return textimgdisto.CosVShift(dst, m.stepRadian, m.ampValue)
}

type distroBlockyFlip7 struct {
	blockWidth  int
	blockHeight int
}

func (m *distroBlockyFlip7) distro(dst *image.Gray) *image.Gray {
	return textimgdisto.BlockyFlip7(dst, m.blockWidth, m.blockHeight)
}

type distroInvert struct {
}

func (m *distroInvert) distro(dst *image.Gray) *image.Gray {
	return textimgdisto.Invert(dst)
}

func parseDistroCommand(cmd string) (m distroMethod, ok bool) {
	aux := strings.Split(cmd, ",")
	auxCnt := len(aux)
	if auxCnt < 1 {
		return
	}
	switch aux[0] {
	case "cosv":
		if auxCnt != 3 {
			log.Print("WARN: expect `cosv,0.07,6.0`; got:", aux)
			return
		}
		stepRadian, err := strconv.ParseFloat(aux[1], 64)
		if nil != err {
			log.Printf("WARN: cannot parse cosv step-radian [%v]: %v", aux[1], err)
			return
		}
		ampValue, err := strconv.ParseFloat(aux[2], 64)
		if nil != err {
			log.Printf("WARN: cannot parse cosv amp-value [%v]: %v", aux[2], err)
			return
		}
		m = &distroCosVShift{
			stepRadian: stepRadian,
			ampValue:   ampValue,
		}
		ok = true
	case "cosh":
		if auxCnt != 3 {
			log.Print("WARN: expect `cosh,0.12,6.0`; got:", aux)
			return
		}
		stepRadian, err := strconv.ParseFloat(aux[1], 64)
		if nil != err {
			log.Printf("WARN: cannot parse cosh step-radian [%v]: %v", aux[1], err)
			return
		}
		ampValue, err := strconv.ParseFloat(aux[2], 64)
		if nil != err {
			log.Printf("WARN: cannot parse cosh amp-value [%v]: %v", aux[2], err)
			return
		}
		m = &distroCosHShift{
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
		m = &distroBlockyFlip7{
			blockWidth:  int(blockWidth),
			blockHeight: int(blockHeight),
		}
		ok = true
	case "inv":
		m = &distroInvert{}
		ok = true
	}
	return
}
