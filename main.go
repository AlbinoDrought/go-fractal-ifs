package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

const A = 0
const B = 1
const C = 2
const D = 3
const E = 4
const F = 5

type Group struct {
	Args        []float32
	Probability float32
}

const TypeRandom = "random"
const TypeFork = "fork"

type IFS struct {
	Rand   *rand.Rand
	Depth  int
	Scale  float32
	Groups []Group
	Type   string
}

type IFSTypeDriver interface {
	Next(system *IFS, lastCall *drawCall) []drawCall
}

type RandomDriver struct{}

func (driver RandomDriver) Next(system *IFS, call *drawCall) (drawCalls []drawCall) {
	group := system.PickGroup()
	nextCall := drawCall{
		x:     (group.Args[A] * call.x) + (group.Args[B] * call.y) + group.Args[E],
		y:     (group.Args[C] * call.x) + (group.Args[D] * call.y) + group.Args[F],
		depth: call.depth - 1,
	}

	if nextCall.depth > 0 {
		drawCalls = append(drawCalls, nextCall)
	}

	return
}

type ForkDriver struct{}

func (driver ForkDriver) Next(system *IFS, call *drawCall) (drawCalls []drawCall) {
	newDepth := call.depth - 1

	if newDepth <= 0 {
		return
	}

	for _, group := range system.Groups {
		nextCall := drawCall{
			x:     (group.Args[A] * call.x) + (group.Args[B] * call.y) + group.Args[E],
			y:     (group.Args[C] * call.x) + (group.Args[D] * call.y) + group.Args[F],
			depth: newDepth,
		}
		drawCalls = append(drawCalls, nextCall)
	}

	return
}

type drawCall struct {
	x     float32
	y     float32
	depth int
}

func (system IFS) PickGroup() (chosen Group) {
	number := system.Rand.Float32()
	tally := float32(0.0)
	for _, chosen = range system.Groups {
		tally += chosen.Probability
		if tally > number {
			return
		}
	}

	return
}

func (system IFS) Draw(img *image.NRGBA, offsetX int, offsetY int) {
	var call drawCall
	var drawCalls []drawCall
	var driver IFSTypeDriver

	if system.Type == TypeRandom {
		driver = RandomDriver{}
	} else if system.Type == TypeFork {
		driver = ForkDriver{}
	} else {
		panic(fmt.Errorf("Unsupported IFS type: %v", system.Type))
	}

	for i := 0; i < 5; i++ {
		drawCalls = append(drawCalls, drawCall{
			x:     0.0,
			y:     0.0,
			depth: system.Depth,
		})
	}

	color := color.NRGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 255,
	}

	for len(drawCalls) > 0 {
		call, drawCalls = drawCalls[0], drawCalls[1:]

		img.Set(int(call.x*system.Scale)+offsetX, int(-call.y*system.Scale)+offsetY, color)

		nextCalls := driver.Next(&system, &call)

		for _, nextCall := range nextCalls {
			drawCalls = append(drawCalls, nextCall)
		}
	}
}

type GeneratedImage struct {
	System  *IFS
	Width   int
	Height  int
	OffsetX int
	OffsetY int
	Name    string
}

func (generatedImage GeneratedImage) Make() *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, generatedImage.Width, generatedImage.Height))

	generatedImage.System.Draw(img, generatedImage.OffsetX, generatedImage.OffsetY)

	return img
}

func (generatedImage GeneratedImage) Save() {
	img := generatedImage.Make()
	f, _ := os.Create(generatedImage.Name + ".png")
	defer f.Close()
	png.Encode(f, img)
}

func main() {
	GeneratedImage{
		Name:    "fern",
		System:  fern(),
		Width:   512,
		Height:  512,
		OffsetX: 256,
		OffsetY: 400,
	}.Save()

	GeneratedImage{
		Name:    "stickfern",
		System:  stickfern(),
		Width:   512,
		Height:  512,
		OffsetX: 256,
		OffsetY: 256,
	}.Save()

	GeneratedImage{
		Name:    "dragon",
		System:  dragon(),
		Width:   512,
		Height:  512,
		OffsetX: 256,
		OffsetY: 300,
	}.Save()

}
