package muggins

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

const NUM_TILES_PER_PLAYER int = 7
const GAME_OVER_SCORE int = 365
const PIPS_SUM int = 168
const PRIV_CUTOFF int = 300
const TABLEAU_SEP string = " | "
const TILE_SEP string = ":"

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

func FindRotation(tableau string, end *End) int {
	lines := strings.Split(tableau, "\n")
	endStr := strconv.Itoa(end.singleValue)
	for idx, line := range lines {
		if strings.Contains(line, endStr) {
			if idx == 0 {
				return 1
			} else if idx < len(lines)-1 {
				if strings.Index(line, endStr) == 0 {
					return 2
				}
				return 0
			}
		}
	}
	return 3
}

func withSpaces(simpleStr string) string {
	lines := strings.Split(simpleStr, "\n")
	totalLineLength := 0
	for _, line := range lines {
		if len(line) > totalLineLength {
			totalLineLength = len(line)
		}
	}
	resultStr := ""
	for idx, line := range lines {
		resultStr += line + strings.Repeat(" ", totalLineLength-len(line))
		if idx != len(lines)-1 {
			resultStr += "\n"
		}
	}
	if len(lines) > totalLineLength {
		hlines := strings.Split(resultStr, "\n")
		resultStr = ""
		for idx, hline := range hlines {
			resultStr += hline + strings.Repeat(" ", len(lines)-totalLineLength)
			if idx != len(hlines)-1 {
				resultStr += "\n"
			}
		}

	} else if len(lines) < totalLineLength {
		resultStr += strings.Repeat("\n"+strings.Repeat(" ", totalLineLength), totalLineLength-len(lines))
	}
	return resultStr
}

func withoutSpaces(squareStr string) string {
	resultStr := ""
	reducedStrList := strings.Split(strings.TrimRight(squareStr, "\n\t "), "\n")
	for idx, line := range reducedStrList {
		resultStr += strings.TrimRight(line, "\t ")
		if idx != len(reducedStrList)-1 {
			resultStr += "\n"
		}
	}
	return resultStr
}

func rotateStringOnce(squareStr string) string {
	matrix := make([][]rune, 0)
	for _, line := range strings.Split(squareStr, "\n") {
		matrix = append(matrix, []rune(line))
	}
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	rotatedStr := ""
	for idx, row := range matrix {
		rotatedStr += string(row)
		if idx != len(matrix)-1 {
			rotatedStr += "\n"
		}
	}
	return rotatedStr
}

func Rotate(tableau string, rotAmt int) string {
	resultTableau := withSpaces(tableau)
	for i := 0; i < (rotAmt % 4); i++ {
		resultTableau = rotateStringOnce(resultTableau)
	}
	return withoutSpaces(resultTableau)
}

func GetMaxSame(hand []Tile) int {
	tilesGrouped := make(map[int][]Tile)
	for _, t := range hand {
		if !t.IsDouble() {
			tilesGrouped[t.left] = append(tilesGrouped[t.left], t)
			tilesGrouped[t.right] = append(tilesGrouped[t.right], t)
		}
	}
	maxSame := 0
	for _, v := range tilesGrouped {
		if len(v) > maxSame {
			maxSame = len(v)
		}
	}
	return maxSame
}
