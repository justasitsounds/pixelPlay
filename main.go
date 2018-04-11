package main

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const CAN_WIDTH float64 = 1024
const CAN_HEIGHT float64 = 768
const X_PIXELS int = 12
const Y_PIXELS int = 8

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, CAN_WIDTH, CAN_HEIGHT),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	points := definePoints(CAN_WIDTH, CAN_HEIGHT, X_PIXELS, Y_PIXELS)

	frames := 0

	for !win.Closed() {
		win.Clear(colornames.Black)
		frames++
		for _, v := range points {
			a := ledpixel(v, colorShader(v, frames, 0.2), 20)
			a.Draw(win)

			b := ledpixel(v, colorShader(v, frames, 0.4), 14)
			b.Draw(win)

			c := ledpixel(v, colorShader(v, frames, 1), 5)
			c.Draw(win)

		}
		win.Update()
	}
}

func definePoints(canvasWidth, canvasHeight float64, columns, rows int) []pixel.Vec {
	numpoints := rows * columns
	pixPerRow := int(canvasHeight) / rows
	pixPerCol := int(canvasWidth) / columns
	points := make([]pixel.Vec, numpoints)
	index := 0

	for y := pixPerRow / 2; y < int(canvasHeight); y += pixPerRow {
		for x := pixPerCol / 2; x < int(canvasWidth); x += pixPerCol {
			points[index] = pixel.V(float64(x), float64(y))
			index++
		}
	}

	return points
}

func colorShader(pos pixel.Vec, frame int, a float64) color.Color {
	r := pos.X / CAN_WIDTH * a
	g := ((math.Sin(float64(frame)/120) / 2) + 0.5) * a
	b := ((math.Cos(float64(frame)/90) / 2) + 0.5) * a
	return pixel.RGB(r, g, b)
}

func ledpixel(pos pixel.Vec, col color.Color, r float64) *imdraw.IMDraw {
	c := imdraw.New(nil)
	c.Color = col
	c.Push(pos)
	c.Circle(r, 0)
	return c
}

func main() {
	pixelgl.Run(run)
}
