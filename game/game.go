package game

import (
	"log"

	"github.com/manleydev/golang-game-framework/input"
	"github.com/manleydev/golang-game-framework/node"
	"github.com/manleydev/golang-game-framework/rendering"
	"github.com/manleydev/golang-game-framework/sdl"
)

func Run(root *node.Node, renderer rendering.Renderer2D, inputHandler input.InputHandler) {
	texture := sdl.NewSdlTexture2D(renderer, "gopher.bmp")

	g := node.NewSprite2D(texture)
	g.Name = "Gopher"
	defer root.AddChild(g)

	g.OnReady(func() error {
		log.Println("Hello world!")
		return nil
	})

	rotation_speed := 0.01
	speed := 0.02
	g.OnUpdate(func() error {

		if inputHandler.IsKeyJustReleased(sdl.SCANCODE_SPACE) {
			rotation_speed = -rotation_speed
		}

		if inputHandler.IsKeyDown(sdl.SCANCODE_D) {
			g.Position.X += speed
		}
		if inputHandler.IsKeyDown(sdl.SCANCODE_A) {
			g.Position.X -= speed
		}
		if inputHandler.IsKeyDown(sdl.SCANCODE_S) {
			g.Position.Y += speed
		}
		if inputHandler.IsKeyDown(sdl.SCANCODE_W) {
			g.Position.Y -= speed
		}

		g.Rotation += rotation_speed
		return nil
	})

	controller := node.NewNode()
	controller.Name = "Controller"
	controller.SetProcessMode(node.ActiveProcessMode)
	defer root.AddChild(&controller)

	controller.OnUpdate(func() error {
		if inputHandler.IsKeyJustPressed(sdl.SCANCODE_ESCAPE) {
			if g.IsProcessing() {
				g.SetProcessMode(node.PausedProcessMode)
			} else {
				g.SetProcessMode(node.DefaultProcessMode)
			}
		}

		if inputHandler.IsKeyJustPressed(sdl.SCANCODE_P) {
			if err := g.Free(); err != nil {
				return err
			}
		}

		return nil
	})

	camera := node.NewCamera2D(&renderer, rendering.NewStandardCamera())
	defer root.AddChild(camera)

	camera.OnReady(func() error {
		return camera.Enable()
	})

}
