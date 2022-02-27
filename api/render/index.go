package handler

import (
	"bufio"
	"bytes"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/anuraghazra/go-phyllotaxis/utils"
	"github.com/fogleman/gg"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func render(canvasSize float64, pointSize float64) *gg.Context {
	var CANVAS_WIDTH float64 = canvasSize
	var CANVAS_HEIGHT float64 = canvasSize
	var N int = int(canvasSize + 200)
	c := gg.NewContext(int(CANVAS_WIDTH), int(CANVAS_HEIGHT))

	var n float64 = utils.RandRange(1, 5)
	var t float64 = utils.RandRange(6, 9)
	c.SetRGB(0, 0, 0)
	c.Clear()

	max := 160
	min := 30
	var rotationAngle float64 = utils.RandRange(min, max)
	for i := 0; i <= N; i++ {
		var a float64 = n * rotationAngle
		var r float64 = t * math.Sqrt(n)

		x := r*math.Cos(a) + CANVAS_WIDTH/2
		y := r*math.Sin(a) + CANVAS_HEIGHT/2

		hue := utils.Normalize(r*math.Cos(a)/360, 1, 0)
		rgb := utils.HSL{hue, .8, .5}.ToRGB()

		c.SetRGB(rgb.R, rgb.G, rgb.B)
		c.DrawCircle(x, y, pointSize)
		n++
		c.Fill()
	}
	return c
}

func Handler(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	canvasSize, _ := strconv.Atoi(query.Get("canvas_size"))
	pointSize, _ := strconv.Atoi(query.Get("point_size"))

	if canvasSize == 0 {
		canvasSize = 400
	}
	if pointSize == 0 {
		pointSize = 5
	}
	var context = render(float64(canvasSize), float64(pointSize))

	var buf bytes.Buffer
	foo := bufio.NewWriter(&buf)

	context.EncodePNG(foo)
	req.Header.Set("Content-Type", `"image/png`)
	w.Write(buf.Bytes())
}
