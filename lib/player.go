package muggins

import (
	"math/rand"

	"golang.org/x/exp/maps"
)

type Player struct {
	ID    string
	hand  []Tile
	score int
}

func (p *Player) DrawTile(r *Round) {
	random_idx := rand.Intn(len(r.reserve))
	p.hand = append(p.hand, r.reserve[random_idx])
	r.reserve = append(r.reserve[:random_idx], r.reserve[random_idx+1:]...)
}

func (p *Player) MakeMove(r *Round) {
	// TODO: Solve the r.line problem: how to identify the non-existence, existence, and emergence of r.line?
	// TODO: Solve the dropping of more than one tiles when they are divisible by 5
	elligibilityMap := make(map[Tile][]int)
	elligibleTiles := p.hand
	if len(r.ends) > 0 {
		for _, v := range p.hand {
			v_el := make([]int, 0)
			for _, end := range r.ends {
				if v.IsPlayable(end.singleValue) {
					v_el = append(v_el, end.singleValue)
				}
			}
			if len(v_el) > 0 {
				elligibilityMap[v] = v_el
			}
		}
		elligibleTiles = maps.Keys(elligibilityMap)
	}
	for len(elligibleTiles) == 0 && len(r.reserve) > 1 {
		p.DrawTile(r)
		lastDrawn := p.hand[len(p.hand)-1]
		el := make([]int, 0)
		for _, end := range r.ends {
			if lastDrawn.IsPlayable(end.singleValue) {
				el = append(el, end.singleValue)
			}
		}
		if len(el) > 0 {
			elligibilityMap[lastDrawn] = el
			elligibleTiles = []Tile{lastDrawn}
		}
	}
	if len(elligibleTiles) == 0 {
		return
	}
	choice := GetChoice(r.game.Strm, elligibleTiles)
	choicePos := -1
	for pos, v := range p.hand {
		if v.IsSame(choice) {
			choicePos = pos
			break
		}
	}
	p.hand = append(p.hand[:choicePos], p.hand[choicePos+1:]...)
	r.tableau = append(r.tableau, choice)
	if len(r.ends) == 0 {
		if choice.IsDouble() {
			r.ends = []End{{singleValue: choice.left, isDouble: true}}
		} else {
			r.ends = []End{
				{singleValue: choice.left, isDouble: false},
				{singleValue: choice.right, isDouble: false},
			}
		}
		return
	}
	elligibleEnds := elligibilityMap[choice]
	choiceEnd := GetChoice(r.game.Strm, elligibleEnds)
	choiceEndPos := -1
	for pos, e := range r.ends {
		if e.singleValue == choiceEnd {
			choiceEndPos = pos
			break
		}
	}
	newEnd := End{singleValue: choice.left, isDouble: choice.IsDouble()}
	if choice.left == choiceEnd {
		newEnd.singleValue = choice.right
	}
	r.ends[choiceEndPos] = newEnd
}
