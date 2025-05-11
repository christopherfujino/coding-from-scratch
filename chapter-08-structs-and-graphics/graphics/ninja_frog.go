package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	//image_color "image/color"
)

// Number of IDLE_FRAME_COUNT in the Ninja Frog animation
const IDLE_FRAME_COUNT = 11
const RUNNING_FRAME_COUNT = 12

const MAX_SPEED = 20
const IDLE_IMAGE = "./assets/ninja_frog/idle-32-32.png"
const RUNNING_IMAGE = "./assets/ninja_frog/run-32-32.png"
const FRAME_WIDTH = 32
const FRAME_HEIGHT = 32

type FrogState uint8

const (
	Idle FrogState = iota
	Running
)

var origin = rl.Vector2{X: 0, Y: 0}

var idleImage *rl.Image
var runningImage *rl.Image

type NinjaFrog struct {
	// Do the textures need to live here, or can they be global?
	IdleTexture    rl.Texture2D
	RunningTexture rl.Texture2D
	FrameIndex     int32
	Position       rl.Vector2
	Speed          int
	State          FrogState
}

// TODO this should be light weight.
func CreateNinjaFrog() *NinjaFrog {
	if idleImage == nil {
		idleImage = rl.LoadImage(IDLE_IMAGE)
	}
	var idleTexture = rl.LoadTextureFromImage(idleImage)
	if runningImage == nil {
		runningImage = rl.LoadImage(RUNNING_IMAGE)
	}
	var runningTexture = rl.LoadTextureFromImage(runningImage)
	// TODO unload the image, if we re-use the texture
	var position = rl.Vector2{X: 10, Y: 10}

	return &NinjaFrog{
		IdleTexture:    idleTexture,
		RunningTexture: runningTexture,
		FrameIndex:     0,
		Position:       position,
		State:          Idle,
		Speed:          0,
	}
}
func (n NinjaFrog) GetTexture() rl.Texture2D {
	switch n.State {
	case Idle:
		return n.IdleTexture
	case Running:
		return n.RunningTexture
	default:
		panic(fmt.Sprintf("unimplemented state %d", n.State))
	}
}

func (n NinjaFrog) GetAnimationRect() rl.Rectangle {
	return rl.Rectangle{
		X:      float32(FRAME_WIDTH * n.FrameIndex),
		Y:      0,
		Width:  FRAME_WIDTH,
		Height: FRAME_HEIGHT,
	}
}

func (n *NinjaFrog) IncrementFrame() {
	n.FrameIndex += 1
	switch n.State {
	case Idle:
		if n.FrameIndex >= IDLE_FRAME_COUNT {
			n.FrameIndex = 0
		}
	case Running:
		if n.FrameIndex >= RUNNING_FRAME_COUNT {
			n.FrameIndex = 0
		}
	}
}

func (n *NinjaFrog) ChangeState(nextState FrogState) {
	if n.State == nextState {
		return
	}
	n.State = nextState
	// Reset animation frame index
	n.FrameIndex = 0
}

func (n *NinjaFrog) Decelerate() bool {
	if n.Speed > 0 {
		n.Speed -= DECELERATION_MAGNITUDE
		if n.Speed <= 0 {
			n.Speed = 0
		}
		return true
	} else if n.Speed < 0 {
		n.Speed += DECELERATION_MAGNITUDE
		if n.Speed >= 0 {
			n.Speed = 0
		}
		return true
	}
	return false
}

const DECELERATION_MAGNITUDE = 1
const ACCELERATION_MAGNITUDE = DECELERATION_MAGNITUDE + 1

func (n *NinjaFrog) Render() {
	if !n.Decelerate() {
		n.ChangeState(Idle)
	}
	if rl.IsKeyDown(rl.KeyL) {
		n.Speed = min(n.Speed+ACCELERATION_MAGNITUDE, MAX_SPEED)
		n.ChangeState(Running)
	}
	if rl.IsKeyDown(rl.KeyH) {
		n.Speed = max(n.Speed-ACCELERATION_MAGNITUDE, -MAX_SPEED)
		n.ChangeState(Running)
	}
	n.Position.X += float32(n.Speed)
	var sourceRect = rl.Rectangle{
		X:      float32(FRAME_WIDTH * n.FrameIndex),
		Y:      0,
		Width:  float32(FRAME_WIDTH),
		Height: float32(FRAME_HEIGHT),
	}

	var destinationRect = rl.Rectangle{
		X:      n.Position.X * scale,
		Y:      n.Position.Y * scale,
		Width:  float32(FRAME_WIDTH * scale),
		Height: float32(FRAME_HEIGHT * scale),
	}

	rl.DrawTexturePro(n.GetTexture(), sourceRect, destinationRect, origin, rotation, rl.White)

	n.IncrementFrame()
}
