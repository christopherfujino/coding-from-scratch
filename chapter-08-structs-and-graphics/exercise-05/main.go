package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const height = 480
const width = 640

func main() {
	rl.InitWindow(width, height, "Title")
	rl.SetTargetFPS(60)

	// Did the user hit escape or close the window?
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		const circleCenterX = 20
		const circleCenterY = 20
		const circleRadius = 10
		var circleColor = rl.Pink
		rl.DrawCircle(circleCenterX, circleCenterY, circleRadius, circleColor)

		const rectangleX = 500
		const rectangleY = 300
		const rectangleWidth = 50
		const rectangleHeight = 75
		var rectangleColor = rl.Blue
		rl.DrawRectangle(rectangleX, rectangleY, rectangleWidth, rectangleHeight, rectangleColor)

		rl.EndDrawing()
	}
}
