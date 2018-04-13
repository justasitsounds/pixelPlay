package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"time"

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
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)

	points := definePoints(CAN_WIDTH, CAN_HEIGHT, X_PIXELS, Y_PIXELS)
	var frames = 0
	var frameCount = 0
	var second = time.Tick(time.Second)
	shaders := []shaderFunc{xGradientShader, colorShader}
	activeShaderIndex := 0

	for !win.Closed() {
		frames++
		frameCount++

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			log.Println("pressed")
			activeShaderIndex++
			if activeShaderIndex == len(shaders) {
				activeShaderIndex = 0
			}
		}

		win.Clear(colornames.Black)
		for _, v := range points {
			b := toBulb(v, 20.0, frames, shaders[activeShaderIndex])

			b.Draw(win)
		}
		win.Update()
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frameCount))
			frameCount = 0
		default:
		}
	}
}

func smoothing(i, limit int) float64 {
	x := (float64(i) / float64(limit))
	return 1 - math.Pow(x, 2)
}

type shaderFunc func(pos pixel.Vec, frame int, a float64) color.Color
type bulb []*imdraw.IMDraw

func toBulb(p pixel.Vec, size float64, time int, shaderFn shaderFunc) bulb {

	steps := 6
	bulbs := make([]*imdraw.IMDraw, steps)
	for i := 1; i <= steps; i++ {
		shader := shaderFn(p, time, 1-smoothing(i, steps))
		circle := ledpixel(p, shader, size*smoothing(i, steps))
		bulbs[i-1] = circle
	}
	return bulbs
}

func (b bulb) Draw(t pixel.Target) {
	for _, v := range b {
		if v != nil {
			v.Draw(t)
		}
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

func xGradientShader(pos pixel.Vec, frame int, a float64) color.Color {
	rate := frame * 3
	r := pos.X / CAN_WIDTH
	g := ((math.Sin(float64(rate)/120) / 2) + 0.5)
	b := ((math.Cos(float64(rate)/90) / 2) + 0.5)
	return pixel.RGB(r, g, b).Mul(pixel.Alpha(a))
}

func colorShader(pos pixel.Vec, frame int, a float64) color.Color {
	var r, g, b float64
	r = 1.0
	g = 0.5
	b = 0.1
	return pixel.RGB(r, g, b).Mul(pixel.Alpha(a))
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
