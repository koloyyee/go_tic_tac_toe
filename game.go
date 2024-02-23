// This is a Tic Tac Toe game for practicing Go.

package main

import (
	"fmt"
	"strings"
)

type Game struct {
	Board  map[string]string
	Player string
	Turn   int
}

func main() {
	var g Game
	g.Player = "O"
	g.Board = map[string]string{
		"1": " ", "2": " ", "3": " ",
		"4": " ", "5": " ", "6": " ",
		"7": " ", "8": " ", "9": " ",
	}
	g.Turn = 9     // set the rounds to 9
	g.PrintBoard() // initialize the game.

	for i := g.Turn; i > 0; i-- {
		if i%2 == 0 {
			g.Player = "X"
		} else {
			g.Player = "O"
		}

		pos := AskPlayer()

		checkedPos, err := g.CheckInput(pos) // validation of the input of position
		if err != nil {
			fmt.Println(err)
			i++ // Keep the loop going because the position has been taken.
		}

		g.Board[checkedPos] = g.Player

		g.PrintBoard()
		winner, isOver := g.CheckWinner()
		if isOver == true && i > 1 {
			fmt.Printf("Winner is %v! \n", strings.TrimRight(winner, "\n"))
			restart := AskRestart()
			if restart == true {
				i = 9
			}
			break
		} else if i == 1 { // 1 because if it is zero it will 10 turns
			fmt.Println("The game has tied!")
			restart := AskRestart()
			if restart == true {
				i = 9
			}
			break
		}
	}
}

// PrintBoard is a method to build a grid for the game
func (g *Game) PrintBoard() {
	fmt.Println(g.Board["1"]+"|", g.Board["2"]+"|", g.Board["3"])
	fmt.Println("-+--+--")
	fmt.Println(g.Board["4"]+"|", g.Board["5"]+"|", g.Board["6"])
	fmt.Println("-+--+--")
	fmt.Println(g.Board["7"]+"|", g.Board["8"]+"|", g.Board["9"])
}

// CheckWinner is a method, check has anyone won the game, the pointer value make sure we are mutating the same board.
func (g *Game) CheckWinner() (winner string, isOver bool) {
	switch {
	case g.Board["7"] == g.Player && g.Board["8"] == g.Player && g.Board["9"] == g.Player:
		return g.Player + "\n", true
	case g.Board["4"] == g.Player && g.Board["5"] == g.Player && g.Board["6"] == g.Player:
		return g.Player + "\n", true
	case g.Board["1"] == g.Player && g.Board["2"] == g.Player && g.Board["3"] == g.Player:
		return g.Player + "\n", true
	case g.Board["7"] == g.Player && g.Board["5"] == g.Player && g.Board["3"] == g.Player:
		return g.Player + "\n", true
	case g.Board["9"] == g.Player && g.Board["5"] == g.Player && g.Board["1"] == g.Player:
		return g.Player + "\n", true
	case g.Board["8"] == g.Player && g.Board["5"] == g.Player && g.Board["2"] == g.Player:
		return g.Player + "\n", true
	case g.Board["7"] == g.Player && g.Board["4"] == g.Player && g.Board["1"] == g.Player:
		return g.Player + "\n", true
	case g.Board["9"] == g.Player && g.Board["6"] == g.Player && g.Board["3"] == g.Player:
		return g.Player + "\n", true
	}

	return "", false
}

// AskPlayer taken an input from the user in console, need to validate the input and check the spot has been taken
func AskPlayer() string {
	fmt.Println("From left to right, top to bottom 1-9 pick your position.")
	var pos string
	fmt.Scanln(&pos)
	return pos
}

// CheckInput checks has the position been take, if it does we will return error message.
func (g *Game) CheckInput(pos string) (string, error) {
	if g.Board[pos] != " " {
		return "", fmt.Errorf("The postion %v has been taken", pos)
	}
	return pos, nil
}

// AskRestart ask the players whether they would like to play another round
func AskRestart() bool {
	var willRestart string
	fmt.Println("Game has ended, would like to play another round? (input 'Y(yes)/N(no)')")
	fmt.Scanln(&willRestart)
	switch strings.ToLower(willRestart) {
	case "y", "yes":
		return true
	case "n", "no":
		return false
	}
	return false
}
