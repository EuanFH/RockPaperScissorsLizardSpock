package main

import (
	"fmt"
	"strings"
)

func main() {
	rockPaperLizard := RockPaperLizard{
		DecisionMap{
			"rock":     {"rock": draw, "paper": loss, "scissors": win, "lizard": win, "spock": loss},
			"paper":    {"rock": win, "paper": draw, "scissors": loss, "lizard": loss, "spock": win},
			"scissors": {"rock": loss, "paper": win, "scissors": draw, "lizard": win, "spock": loss},
			"lizard":   {"rock": loss, "paper": win, "scissors": loss, "lizard": draw, "spock": win},
			"spock":    {"rock": win, "paper": loss, "scissors": win, "lizard": loss, "spock": draw},
		},
	}
	fmt.Println("Welcome to Rock Paper Scissors Lizard Spock")
	for {
		aigesture := rockPaperLizard.RandomGesture()
		fmt.Println("Please enter a gesture or exit to leave the game")
		userInput := getUserInput()
		if userInput == "exit" {
			return
		}
		if !rockPaperLizard.validGesture(userInput) || !rockPaperLizard.validGesture(aigesture) {
			fmt.Println("invalid play")
			continue
		}
		gameOutcome := rockPaperLizard.GameResult(userInput, aigesture)
		fmt.Println("You Played: ", userInput, "\nAi Played: ", aigesture)
		switch gameOutcome {
		case loss:
			fmt.Println("loss")
		case draw:
			fmt.Println("draw")
		case win:
			fmt.Println("win")
		default:
			fmt.Println("unknown")
		}
	}
}

func getUserInput() string {
	var input string
	if _, err := fmt.Scanf("%s\n", &input); err != nil {
		return ""
	}
	return strings.ToLower(input)
}
