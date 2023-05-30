package muggins

import (
	"fmt"
	"log"
	"strconv"

	"golang.org/x/exp/constraints"
)

const NUM_TILES_PER_PLAYER int = 7
const GAME_OVER_SCORE int = 365
const PIPS_SUM int = 168
const PRIV_CUTOFF int = 300

func GetMin[T constraints.Ordered](x T, y T) T {
	if x <= y {
		return x
	}
	return y
}

func GetChoice[T any](stream Stream, arr []T) T {
	for pos, v := range arr {
		ftmMessage := fmt.Sprintf("%d: Play %v\n", pos, v)
		errW := stream.WriteMessage(1, []byte(ftmMessage))
		if errW != nil {
			log.Fatal(errW)
		}
	}
	stream.WriteMessage(1, []byte("Choose >> "))
	_, input, errR := stream.ReadMessage()
	if errR != nil {
		log.Fatal(errR)
	}
	intInput, errC := strconv.Atoi(string(input))
	for errC != nil || intInput < 0 || intInput >= len(arr) {
		stream.WriteMessage(1, []byte("Invalid input!"))
		stream.WriteMessage(1, []byte("Choose >> "))
		_, input, errR = stream.ReadMessage()
		if errR != nil {
			log.Fatal(errR)
		}
		intInput, errC = strconv.Atoi(string(input))
	}
	return arr[intInput]
}

func Has[T comparable](arr []T, elem T) bool {
	for _, item := range arr {
		if elem == item {
			return true
		}
	}
	return false
}
