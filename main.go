package main

import (
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/maiconspas/go-space-inv/physics"
	"github.com/maiconspas/go-space-inv/utils"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var fps string = strconv.FormatFloat(ebiten.ActualFPS(), 'f', 0, 64)
	ebitenutil.DebugPrint(screen, fps)
	ebitenutil.DebugPrintAt(screen, strconv.FormatInt(int64(physics.ObjectCount), 10), 0, 20)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	physics.StartPhysicSystem()
	physics.CreatePhysicObject(utils.Vector2d{X: 100, Y: 100})
	physics.CreatePhysicObject(utils.Vector2d{X: 50, Y: 100})
	physics.CreatePhysicObject(utils.Vector2d{X: 0, Y: 100})
	physics.CreatePhysicObject(utils.Vector2d{X: 200, Y: 100})
	physics.CreatePhysicObject(utils.Vector2d{X: 150, Y: 100})
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
