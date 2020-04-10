package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/gosuri/uilive"
)

const (
	Refresh_Rate = 150
	Board_Width  = 80
	Board_Height = 20
	Board_Border = "#"
	Board_Tail   = "$"
	Board_Head   = "*"
	Board_fruit  = "@"
	Board_Espace = " "

	SNAKE_STILL     = 0
	DIRECTION_UP    = 1
	DIRECTION_DOWN  = 2
	DIRECTION_LEFT  = 3
	DIRECTION_RIGHT = 4
)

type snake struct {
	body             []snakeBodyPart
	currentDirection int32
	newDirection     int32
}

type snakeBodyPart struct {
	x int
	y int
}

type fruit struct {
	x int
	y int
}

func (s *snake) Switch(direction int32) {
	if len(s.body) > 1 {
		if (s.currentDirection == DIRECTION_UP && direction != DIRECTION_DOWN) ||
			(s.currentDirection == DIRECTION_DOWN && direction != DIRECTION_UP) ||
			(s.currentDirection == DIRECTION_RIGHT && direction != DIRECTION_LEFT) ||
			(s.currentDirection == DIRECTION_LEFT && direction != DIRECTION_RIGHT) {
			s.newDirection = direction
		}
	} else {
		s.newDirection = direction
	}
}

type Board struct {
	snake *snake
	fruit *fruit
}

func (b *Board) init(s *snake) {
	b.snake = s
}

func (s *snake) collisionWithfruit(f fruit) bool {
	for _, body := range s.body {
		if body.x == f.x && body.y == f.y {
			return true
		}
	}
	return false
}

func (s *snake) collisionWithHead(head snakeBodyPart) bool {
	for _, body := range s.body {
		if body.x == head.x && body.y == head.y {
			return true
		}
	}
	return false
}

func (b *Board) moveSnake() {
	if b.snake.newDirection > SNAKE_STILL {
		b.snake.currentDirection = b.snake.newDirection
		b.snake.newDirection = SNAKE_STILL
	}
	if len(b.snake.body) > 0 && b.snake.currentDirection > SNAKE_STILL {
		//Armo nueva cabeza
		var oldHead = b.snake.body[0]
		var head snakeBodyPart
		switch b.snake.currentDirection {
		case DIRECTION_UP:
			head = snakeBodyPart{x: oldHead.x - 1, y: oldHead.y}
			break
		case DIRECTION_DOWN:
			head = snakeBodyPart{x: oldHead.x + 1, y: oldHead.y}
			break
		case DIRECTION_RIGHT:
			head = snakeBodyPart{x: oldHead.x, y: oldHead.y + 1}
			break
		case DIRECTION_LEFT:
			head = snakeBodyPart{x: oldHead.x, y: oldHead.y - 1}
			break
		}
		//Chequeo si choco con el muro
		if head.x < 1 || head.x >= Board_Height || head.y < 1 || head.y >= Board_Width {
			log.Fatalln("Game over!")
			os.Exit(1)
		}
		if b.snake.collisionWithHead(head) {
			log.Fatalln("Game over!")
			os.Exit(1)
		}
		//Armo nuevo cuerpo con nueva cabeza
		b.snake.body = append([]snakeBodyPart{head}, b.snake.body...)
		//Chequeo si comio fruta
		if b.fruit.x == head.x && b.fruit.y == head.y {
			b.addfruit()
		} else {
			//Borro ultimo
			b.snake.body = b.snake.body[:len(b.snake.body)-1]
		}
	}
}

func (b *Board) addfruit() {
	var f fruit
	for {
		f = fruit{x: rand.Intn(Board_Height-2) + 1, y: rand.Intn(Board_Width-2) + 1}
		if !b.snake.collisionWithfruit(f) {
			break
		}
	}
	b.fruit = &f
}

func (b *Board) buildScreen() string {
	//Armo tablero
	var cells [Board_Height][Board_Width]string
	for row := 0; row < Board_Height; row++ {
		for col := 0; col < Board_Width; col++ {
			if col == 0 || col == Board_Width-1 {
				cells[row][col] = Board_Border
			} else if row == 0 || row == Board_Height-1 {
				cells[row][col] = Board_Border
			} else {
				cells[row][col] = Board_Espace
			}
		}
	}

	//Pego vibora
	for _, body := range b.snake.body {
		cells[body.x][body.y] = Board_Head
	}

	//Pego fruta
	cells[b.fruit.x][b.fruit.y] = Board_fruit

	var pantalla = ""
	//Imprimo tablero
	for i, _ := range cells {
		for j, _ := range cells[i] {
			var cel = cells[i][j]
			pantalla += cel
		}
		pantalla += "\n"
	}
	pantalla += "Snake length: " + strconv.Itoa(len(b.snake.body)) + "\n"
	return pantalla
}

func printGame(b *Board) {
	writer := uilive.New()
	writer.Start()
	defer writer.Stop() // flush and sto
	for {
		b.moveSnake()
		var screen = b.buildScreen()
		fmt.Fprintf(writer, screen)
		time.Sleep(Refresh_Rate * time.Millisecond)
	}
}

func main() {
	fmt.Println("Snake Game")
	rand.Seed(time.Now().UnixNano())

	var board = Board{}

	var sbp []snakeBodyPart
	sbp = append(sbp, snakeBodyPart{x: 1, y: 1})

	var snake = snake{body: sbp, currentDirection: SNAKE_STILL}

	board.init(&snake)

	board.addfruit()

	go printGame(&board)

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	fmt.Println("Press ESC to quit")
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		} else if key == keyboard.KeyArrowUp {
			snake.Switch(DIRECTION_UP)
		} else if key == keyboard.KeyArrowDown {
			snake.Switch(DIRECTION_DOWN)
		} else if key == keyboard.KeyArrowLeft {
			snake.Switch(DIRECTION_LEFT)
		} else if key == keyboard.KeyArrowRight {
			snake.Switch(DIRECTION_RIGHT)
		} else if key == keyboard.KeyEsc {
			break
		}
	}

}
