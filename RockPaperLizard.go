package main

import (
	"math/rand"
	"reflect"
)

type GameOutcome int
type DecisionMap map[string]map[string]GameOutcome

//this is an enum
//unknown is 0 which is also the default type for an int that doesn't exist
const (
	unknown GameOutcome = iota
	loss
	draw
	win
)

type RockPaperLizard struct {
	decisionMap DecisionMap
}

func (rpl *RockPaperLizard) validGesture(gesture string) bool {
	_, exist := rpl.decisionMap[gesture]
	return exist
}

//if the gesture doesn't exist in the map it will return 0 which is the unknown value in the GameOutcome type
func (rpl *RockPaperLizard) GameResult(player1 string, player2 string) GameOutcome {
	return rpl.decisionMap[player1][player2]
}

func (rpl *RockPaperLizard) RandomGesture() string {
	keys := reflect.ValueOf(rpl.decisionMap).MapKeys()
	return keys[rand.Intn(len(keys))].String()
}
