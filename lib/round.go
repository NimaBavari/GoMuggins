package muggins

import (
	"fmt"
	"math/rand"
	"strings"
)

type End struct {
	singleValue int
	isDouble    bool
}

type Round struct {
	game     *Game
	line     *Tile
	tableau  string
	ends     []End
	boneyard []Tile
}

func (r Round) GetEndsSum() int {
	s := 0
	for _, v := range r.ends {
		d := 1
		if v.isDouble {
			d = 2
		}
		s += v.singleValue * d
	}
	return s
}

func (r Round) GetTableauSum() int {
	s := 0
	for _, v := range strings.Split(r.tableau, TABLEAU_SEP) {
		s += FromString(v).Face()
	}
	return s
}

func (r Round) GetBoneyardSum() int {
	s := 0
	for _, v := range r.boneyard {
		s += v.Face()
	}
	return s
}

func (r *Round) AddToTableau(tl Tile, e *End) {
	if e == nil {
		r.tableau = fmt.Sprintf("%d%s%d", tl.left, TILE_SEP, tl.right)
		return
	}
	rotationAmt := FindRotation(r.tableau, e)
	tempTableau := Rotate(r.tableau, rotationAmt)
	inner := e.singleValue
	outer := tl.right
	if tl.right == inner {
		outer = tl.left
	}
	ad := fmt.Sprintf("%s %s %d%s%d", tempTableau, TABLEAU_SEP, inner, TILE_SEP, outer)
	r.tableau = Rotate(ad, -rotationAmt)
}

func (r *Round) Setup(g *Game) {
	r.boneyard = []Tile{
		{0, 0}, {0, 1}, {0, 2}, {0, 3},
		{0, 4}, {0, 5}, {0, 6}, {1, 1},
		{1, 2}, {1, 3}, {1, 4}, {1, 5},
		{1, 6}, {2, 2}, {2, 3}, {2, 4},
		{2, 5}, {2, 6}, {3, 3}, {3, 4},
		{3, 5}, {3, 6}, {4, 4}, {4, 5},
		{4, 6}, {5, 5}, {5, 6}, {6, 6},
	}
	players := make([]Player, 0)
	for _, player := range g.Players {
		h := r.Distribute()
		player.hand = h
		players = append(players, player)
	}
	g.Players = players
	r.game = g
}

func (r *Round) Distribute() []Tile {
	// TODO: Add support for hand contents criteria
	hand := make([]Tile, NUM_TILES_PER_PLAYER)
	for i := 0; i < NUM_TILES_PER_PLAYER; i++ {
		random_idx := rand.Intn(len(r.boneyard))
		hand[i] = r.boneyard[random_idx]
		r.boneyard = append(r.boneyard[:random_idx], r.boneyard[random_idx+1:]...)
	}
	return hand
}

func (r *Round) CheckSetLine() {
	if r.line != nil {
		return
	}
	tiles_down_str := strings.Split(r.tableau, TABLEAU_SEP)
	for idx, v := range tiles_down_str {
		currentTile := FromString(v)
		if currentTile.IsDouble() && idx != 0 && idx != len(tiles_down_str)-1 {
			r.line = &currentTile
			return
		}
	}
}
