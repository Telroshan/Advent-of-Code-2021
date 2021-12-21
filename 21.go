package main

import "fmt"

var dieValue int = 1
var diceRolled int = 0

func rollDie() int {
	value := dieValue
	dieValue = dieValue + 1
	if dieValue == 101 {
		dieValue = 1
	}
	diceRolled++
	return value
}

func clampPosition(pos int) int {
	value := pos % 10
	if value == 0 {
		return 10
	}
	return value
}

func main() {
	player1Position := 1
	player2Position := 10
	player1Score := 0
	player2Score := 0
	turn := 0
	for player1Score < 1000 && player2Score < 1000 {
		movement := rollDie() + rollDie() + rollDie()
		if turn == 0 {
			player1Position = clampPosition(player1Position + movement)
			player1Score += player1Position
		} else {
			player2Position = clampPosition(player2Position + movement)
			player2Score += player2Position
		}

		turn = (turn + 1) % 2
	}

	if player1Score > player2Score {
		fmt.Println("Player 1 wins :", player1Score,"=>",player2Score,"*",diceRolled,"=",player2Score*diceRolled)
	} else {
		fmt.Println("Player 2 wins :", player2Score,"=>",player1Score,"*",diceRolled,"=",player1Score*diceRolled)
	}
}