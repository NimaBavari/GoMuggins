package muggins

import (
	"math/rand"
)

type Pairing struct {
	end int
	tl  Tile
}

type Player struct {
	ID    string
	hand  []Tile
	score int
}

func (p *Player) DrawTile(r *Round) {
	random_idx := rand.Intn(len(r.boneyard))
	p.hand = append(p.hand, r.boneyard[random_idx])
	r.boneyard = append(r.boneyard[:random_idx], r.boneyard[random_idx+1:]...)
}

func (p *Player) MakeMove(r *Round) {
	// TODO: Solve the dropping of more than one tiles when they are divisible by 5
	// TODO: What happens if len(r.ends) == 1?
	if len(r.ends) == 0 {
		firstTile := GetChoice(r.game.Strm, p.hand)
		if firstTile.IsDouble() {
			r.ends = []End{{singleValue: firstTile.left, isDouble: true, isDeveloped: true}}
		} else {
			r.ends = []End{
				{singleValue: firstTile.left, isDeveloped: true},
				{singleValue: firstTile.right, isDeveloped: true},
			}
		}
		r.AddToTableau(firstTile, nil)
		return
	}
	elligibleTiles := make([]Tile, 0)
	pairings := make([]Pairing, 0)
	for _, v := range p.hand {
		for _, end := range r.ends {
			if v.IsPlayable(end.singleValue) {
				pairings = append(pairings, Pairing{end.singleValue, v})
				if !Has(elligibleTiles, v) {
					elligibleTiles = append(elligibleTiles, v)
				}
			}
		}
	}
	for len(elligibleTiles) == 0 && len(r.boneyard) > 1 {
		p.DrawTile(r)
		lastDrawn := p.hand[len(p.hand)-1]
		for _, end := range r.ends {
			if lastDrawn.IsPlayable(end.singleValue) {
				pairings = append(pairings, Pairing{end.singleValue, lastDrawn})
				if !Has(elligibleTiles, lastDrawn) {
					elligibleTiles = append(elligibleTiles, lastDrawn)
				}
			}
		}
	}
	if len(elligibleTiles) == 0 {
		return
	}
	choiceCombo := GetChoice(r.game.Strm, pairings)
	endPos, tilePos := -1, -1
	for pos, e := range r.ends {
		if e.singleValue == choiceCombo.end {
			endPos = pos
			break
		}
	}
	for pos, v := range p.hand {
		if v.IsSame(choiceCombo.tl) {
			tilePos = pos
			break
		}
	}
	p.hand = append(p.hand[:tilePos], p.hand[tilePos+1:]...)
	newEnd := End{singleValue: choiceCombo.tl.left, isDouble: choiceCombo.tl.IsDouble(), isDeveloped: true}
	if choiceCombo.tl.left == choiceCombo.end {
		newEnd.singleValue = choiceCombo.tl.right
	}
	r.ends[endPos] = newEnd
	r.AddToTableau(choiceCombo.tl, &newEnd)
	r.CheckSetLine()
}
