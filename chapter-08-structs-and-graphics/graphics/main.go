package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const size int32 = 32
const width = 1600
const height = 960
const fps = 30

const scale = 3

const rotation = 0

type Element interface {
	Render()
}

func main() {
	rl.InitWindow(width, height, "Ninja Frog")
	rl.SetTargetFPS(fps)

	var elements = []Element{
		CreateNinjaFrog(),
	}
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		for _, element := range elements {
			element.Render()
		}
		rl.EndDrawing()
	}
}
