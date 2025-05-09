package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	//image_color "image/color"
)

// Number of frames in the Ninja Frog animation
const frames = 11

var image *rl.Image
var texture rl.Texture2D

const MAX_SPEED = 20
const IMAGE_PATH = "./tmp/idle-32-32.png"

var origin = rl.Vector2{X: 0, Y: 0}

func CreateNinjaFrog() *NinjaFrog {
	if image == nil {
		image = rl.LoadImage(IMAGE_PATH)
		texture = rl.LoadTextureFromImage(image)
	}
	var frameWidth = texture.Width / frames
	var frameHeight = texture.Height
	var position = rl.Vector2{X: 10, Y: 10}

	return &NinjaFrog{
		Texture:     texture,
		FrameWidth:  frameWidth,
		FrameHeight: frameHeight,
		FrameIndex:  0,
		Position:    position,
	}
}

type NinjaFrog struct {
	Texture     rl.Texture2D
	FrameWidth  int32
	FrameHeight int32
	FrameIndex  int32
	Position    rl.Vector2
	Speed       int
}

const DECELERATION_MAGNITUDE = 3
const ACCELERATION_MAGNITUDE = DECELERATION_MAGNITUDE + 1

func (n *NinjaFrog) Render() {
	if n.Speed > 0 {
		n.Speed -= DECELERATION_MAGNITUDE
		if n.Speed <= 0 {
			n.Speed = 0
		}
	} else if n.Speed < 0 {
		n.Speed += DECELERATION_MAGNITUDE
		if n.Speed >= 0 {
			n.Speed = 0
		}
	}
	if rl.IsKeyDown(rl.KeyL) {
		n.Speed = min(n.Speed + ACCELERATION_MAGNITUDE, MAX_SPEED)
	}
	if rl.IsKeyDown(rl.KeyH) {
		n.Speed = max(n.Speed - ACCELERATION_MAGNITUDE, -MAX_SPEED)
	}
	n.Position.X += float32(n.Speed)
	fmt.Printf("Speed: %v\tPosition: %v\n", n.Speed, n.Position.X)
	var sourceRect = rl.Rectangle{
		X:      float32(n.FrameWidth * n.FrameIndex),
		Y:      0,
		Width:  float32(n.FrameWidth),
		Height: float32(n.FrameHeight),
	}

	var destinationRect = rl.Rectangle{
		// Should this use position?!
		X:      n.Position.X * scale,
		Y:      n.Position.Y * scale,
		Width:  float32(n.FrameWidth * scale),
		Height: float32(n.FrameHeight * scale),
	}

	rl.DrawTexturePro(n.Texture, sourceRect, destinationRect, origin, rotation, rl.White)
	rl.EndDrawing()

	n.FrameIndex += 1
	if n.FrameIndex >= frames {
		n.FrameIndex = 0
	}
}
