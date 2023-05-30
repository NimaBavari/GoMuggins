package muggins

type Game struct {
	NumPlayers int
	Players    []Player
	Strm       Stream
}

func (g *Game) AddPlayer(player Player) {
	g.Players = append(g.Players, player)
}

func (g *Game) FindPlayerWithTile(t Tile) int {
	for pos, player := range g.Players {
		if Has(player.hand, t) {
			return pos
		}
	}
	return -1
}

func (g *Game) DecideFirstPlayer(r Round) int {
	starterTile := Tile{2, 3}
	firstPlayerPos := g.FindPlayerWithTile(starterTile)
	for firstPlayerPos == -1 && len(r.boneyard) > 2 {
		for _, player := range g.Players {
			player.DrawTile(&r)
		}
		firstPlayerPos = g.FindPlayerWithTile(starterTile)
	}
	return firstPlayerPos
}

func (g Game) GetLeadingScore() int {
	leadingScore := 0
	for _, player := range g.Players {
		if player.score > leadingScore {
			leadingScore = player.score
		}
	}
	return leadingScore
}

func (g *Game) Play() {
	gameOver := false
	roundWinnerIdx := -1
	roundIdx := 0
	for !gameOver {
		round := Round{}
		round.Setup(g)
		roundStarterIdx := roundWinnerIdx
		if roundIdx == 0 {
			roundStarterIdx = g.DecideFirstPlayer(round)
		}
		roundEnd := false
		iteration := 0
		for !roundEnd {
			for idx, player := range g.Players {
				if iteration == 0 && idx < roundStarterIdx {
					continue
				}
				player.MakeMove(&round)
				if round.GetEndsSum()%5 == 0 {
					player.score += round.GetEndsSum()
				}
				if len(player.hand) == 0 {
					if player.score < PRIV_CUTOFF {
						otherHandsTotal := PIPS_SUM - (round.GetTableauSum() + round.GetBoneyardSum())
						player.score = GetMin(player.score+otherHandsTotal, PRIV_CUTOFF)
					}
					roundEnd = true
					roundWinnerIdx = idx
					break
				}
			}
			iteration++
		}
		roundIdx++
		if g.GetLeadingScore() >= GAME_OVER_SCORE {
			gameOver = true
		}
	}
	g.Strm.Close()
}
