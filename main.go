package main

import (
	"fmt"
	"image/color"
	"math"
	"reflect"

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
	steps := 6
	size := 20.0

	for !win.Closed() {
		win.Clear(colornames.Black)
		frames++
		for _, v := range points {

			for i := 1; i <= steps; i++ {
				shader := colorShader(v, frames, 1-smoothing(i, steps))
				circle := ledpixel(v, shader, size*smoothing(i, steps))
				circle.Draw(win)
			}
		}
		win.Update()
	}
}

func smoothing(i, limit int) float64 {
	x := (float64(i) / float64(limit)) + 0.01 //0.0 -> 1.0
	return 1 - math.Pow(x, 2)
}

type bulb []*imdraw.IMDraw

func toBulb(p pixel.Vec, size float64, time int) bulb {

	//given an iterator and a limit, return on 1-x^2
	smoothing := func(i, limit int) float64 {
		x := float64(i) / float64(limit) //0.0 -> 1.0
		return 1 - math.Pow(x, 2)
	}

	steps := 4
	bulbs := make([]*imdraw.IMDraw, steps+1)
	for i := 1; i <= steps; i++ {
		shader := colorShader(p, time, 1-smoothing(i, steps))
		circle := ledpixel(p, shader, size*smoothing(i, steps))
		bulbs[i] = circle
	}
	return bulbs
}

func (b bulb) Draw(t pixel.Target) {
	for _, v := range b {
		fmt.Println(reflect.TypeOf(v))
		// v.Draw(t)
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

type shaderFunc func(pos pixel.Vec, frame int, a float64) color.Color

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
