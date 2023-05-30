package muggins

import "math/rand"

type End struct {
	singleValue int
	isDouble    bool
}

type Round struct {
	game    *Game
	line    Tile
	tableau []Tile
	ends    []End
	reserve []Tile
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
	for _, v := range r.tableau {
		s += v.Face()
	}
	return s
}

func (r Round) GetReserveSum() int {
	s := 0
	for _, v := range r.reserve {
		s += v.Face()
	}
	return s
}

func (r *Round) Setup(g *Game) {
	r.reserve = []Tile{
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
		random_idx := rand.Intn(len(r.reserve))
		hand[i] = r.reserve[random_idx]
		r.reserve = append(r.reserve[:random_idx], r.reserve[random_idx+1:]...)
	}
	return hand
}
