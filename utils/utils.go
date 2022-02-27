package utils

import "math/rand"

type RGB struct {
	R, G, B float64
}

type HSL struct {
	H, S, L float64
}

func HueToRGB(v1, v2, h float64) float64 {
	if h < 0 {
		h += 1
	}
	if h > 1 {
		h -= 1
	}
	switch {
	case 6*h < 1:
		return (v1 + (v2-v1)*6*h)
	case 2*h < 1:
		return v2
	case 3*h < 2:
		return v1 + (v2-v1)*((2.0/3.0)-h)*6
	}
	return v1
}

func (c HSL) ToRGB() RGB {
	h := c.H
	s := c.S
	l := c.L

	if s == 0 {
		// it's gray
		return RGB{l, l, l}
	}

	var v1, v2 float64
	if l < 0.5 {
		v2 = l * (1 + s)
	} else {
		v2 = (l + s) - (s * l)
	}

	v1 = 2*l - v2

	r := HueToRGB(v1, v2, h+(1.0/3.0))
	g := HueToRGB(v1, v2, h)
	b := HueToRGB(v1, v2, h-(1.0/3.0))

	return RGB{r, g, b}
}

func Normalize(value float64, max float64, min float64) float64 {
	return (value - min) / (max - min)
}

func RandRange(min int, max int) float64 {
	return float64(rand.Intn(max-min) + min)
}