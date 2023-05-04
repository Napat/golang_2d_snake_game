package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

const ()

type Game struct {
	snake         *Snake
	food          *Food
	score         int
	gameOver      bool
	updateCounter int // value increase every Update and update screen when updateCounter == speed
	speed         int // the lower spped value, the higher snake speed
}

func (g *Game) Update() error {
	if g.gameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.restart()
		}
		return nil
	}

	g.updateCounter++
	if !g.shouldUpdate() {
		return nil
	}

	// reset counter
	g.updateCounter = 0

	// Update the snake's position
	g.snake.Move()

	// Check user input
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.snake.IsDirectionVertical() {
		g.snake.Direction = DIRECTION_LEFT
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) && g.snake.IsDirectionVertical() {
		g.snake.Direction = DIRECTION_RIGHT
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) && g.snake.IsDirectionHorizontal() {
		g.snake.Direction = DIRECTION_UP
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) && g.snake.IsDirectionHorizontal() {
		g.snake.Direction = DIRECTION_DOWN
	}

	// Check gameOver: head
	head := g.snake.Body[0]
	if head.X < 0 || head.Y < 0 || head.X >= SCREEN_WIDTH/TILE_SIZE || head.Y >= SCREEN_HEIGTH/TILE_SIZE {
		g.gameOver = true
		g.speed = SPEED_INIT
	}
	// Check gameOver: body
	for _, part := range g.snake.Body[1:] {
		if head.X == part.X && head.Y == part.Y {
			g.gameOver = true
			g.speed = SPEED_INIT
		}
	}

	// Check eat food
	if head.X == g.food.Position.X && head.Y == g.food.Position.Y {
		g.score += 1
		g.snake.GrowCounter += 1
		g.food = NewFood()

		// Increase speed(Decrease speed value), with a lower limit(max speed)
		if g.speed > SPEED_MAX {
			g.speed--
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw background
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// Draw snake
	for _, p := range g.snake.Body {
		vector.DrawFilledRect(screen, float32(p.X*TILE_SIZE), float32(p.Y*TILE_SIZE), float32(TILE_SIZE), float32(TILE_SIZE), color.RGBA{0, 255, 0, 255}, false)
	}

	// Draw food
	vector.DrawFilledRect(screen, float32(g.food.Position.X*TILE_SIZE), float32(g.food.Position.Y*TILE_SIZE), float32(TILE_SIZE), float32(TILE_SIZE), color.RGBA{255, 0, 0, 255}, false)

	// Create a font.Face
	face := basicfont.Face7x13

	// Draw game over text
	if g.gameOver {
		text.Draw(screen, "Game Over", face, SCREEN_WIDTH/2-40, SCREEN_HEIGTH/2, color.White)
		text.Draw(screen, "Press 'R' to restart", face, SCREEN_WIDTH/2-60, SCREEN_HEIGTH/2+16, color.White)
	}

	// Draw score
	scoreText := fmt.Sprintf("Score: %d", g.score)
	text.Draw(screen, scoreText, face, 5, SCREEN_HEIGTH-5, color.White)
}

// Layout return the fame loginal screen size.
// The screen is automatically scaled.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHight int) {
	return SCREEN_WIDTH, SCREEN_HEIGTH
}

func (g *Game) restart() {
	g.snake = NewSnake()
	g.score = 0
	g.gameOver = false
	g.food = NewFood()
	g.speed = SPEED_INIT
}

func (g *Game) shouldUpdate() bool {
	return g.updateCounter >= g.speed
}
