package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

const (
	Refresh_Rate = 150
	Board_Width  = 80
	Board_Height = 20
	Board_Border = "#"
	Board_Tail   = "$"
	Board_Head   = "*"
	Board_Fruit  = "$"
	Board_Espace = " "

	DIRECTION_UP    = 1
	DIRECTION_DOWN  = 2
	DIRECTION_LEFT  = 3
	DIRECTION_RIGHT = 4
	SNAKE_STILL     = 0
)

type Snake struct {
	body      []SnakeBodyPart
	Direction int32
}

type SnakeBodyPart struct {
	x int
	y int
}

type Fruit struct {
	x int
	y int
}

func (s *Snake) Switch(direction int32) {
	if len(s.body) > 1 {
		if (s.Direction == DIRECTION_UP && direction != DIRECTION_DOWN) ||
			(s.Direction == DIRECTION_DOWN && direction != DIRECTION_UP) ||
			(s.Direction == DIRECTION_RIGHT && direction != DIRECTION_LEFT) ||
			(s.Direction == DIRECTION_LEFT && direction != DIRECTION_RIGHT) {
			s.Direction = direction
		}
	} else {
		s.Direction = direction
	}
}

type Board struct {
	Snake *Snake
	Fruit *Fruit
}

func (b *Board) Init(s *Snake) {
	b.Snake = s
}

func (s *Snake) collisionWithFruit(f Fruit) bool {
	for _, body := range s.body {
		if body.x == f.x && body.y == f.y {
			return true
		}
	}
	return false
}

func (s *Snake) collisionWithHead(head SnakeBodyPart) bool {
	for _, body := range s.body {
		if body.x == head.x && body.y == head.y {
			return true
		}
	}
	return false
}

func (b *Board) MoveSnake() {
	if len(b.Snake.body) > 0 && b.Snake.Direction != SNAKE_STILL {
		//Armo nueva cabeza
		var oldHead = b.Snake.body[0]
		var head SnakeBodyPart
		switch b.Snake.Direction {
		case DIRECTION_UP:
			head = SnakeBodyPart{x: oldHead.x - 1, y: oldHead.y}
			break
		case DIRECTION_DOWN:
			head = SnakeBodyPart{x: oldHead.x + 1, y: oldHead.y}
			break
		case DIRECTION_RIGHT:
			head = SnakeBodyPart{x: oldHead.x, y: oldHead.y + 1}
			break
		case DIRECTION_LEFT:
			head = SnakeBodyPart{x: oldHead.x, y: oldHead.y - 1}
			break
		}
		//Chequeo si choco con el muro
		if head.x < 1 || head.x >= Board_Height || head.y < 1 || head.y >= Board_Width {
			log.Fatalln("Perdiste!")
			os.Exit(1)
		}
		if b.Snake.collisionWithHead(head) {
			log.Fatalln("Perdiste!")
			os.Exit(1)
		}
		//Armo nuevo cuerpo con nueva cabeza
		b.Snake.body = append([]SnakeBodyPart{head}, b.Snake.body...)
		//Chequeo si comio fruta
		if b.Fruit.x == head.x && b.Fruit.y == head.y {
			b.AddFruit()
		} else {
			//Borro ultimo
			b.Snake.body = b.Snake.body[:len(b.Snake.body)-1]
		}
	}
}

func (b *Board) AddFruit() {
	var fruit Fruit
	for {
		fruit = Fruit{x: rand.Intn(Board_Height-2) + 1, y: rand.Intn(Board_Width-2) + 1}
		if !b.Snake.collisionWithFruit(fruit) {
			break
		}
	}
	b.Fruit = &fruit
}

func (b *Board) Print() {
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
	for _, body := range b.Snake.body {
		cells[body.x][body.y] = Board_Head
	}

	//Pego fruta
	cells[b.Fruit.x][b.Fruit.y] = Board_Fruit

	//Imprimo tablero
	for i, _ := range cells {
		for j, _ := range cells[i] {
			var cel = cells[i][j]
			fmt.Print(cel)
		}
		fmt.Println()
	}

}

func printGame(b *Board) {
	for {
		CallClear()
		b.MoveSnake()
		b.Print()
		time.Sleep(Refresh_Rate * time.Millisecond)
	}
}

func main() {
	fmt.Println("Inicio de Snake Game")

	var board = Board{}

	var snakeBodyPart []SnakeBodyPart
	snakeBodyPart = append(snakeBodyPart, SnakeBodyPart{x: 1, y: 1})

	var snake = Snake{body: snakeBodyPart, Direction: SNAKE_STILL}

	board.Init(&snake)

	board.AddFruit()

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
