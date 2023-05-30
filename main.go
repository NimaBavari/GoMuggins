package main

import (
	"log"
	"net/http"

	muggins "github.com/NimaBavari/GoMuggins/lib"
	"github.com/gorilla/websocket"
)

func NewGame(numPlayers int, stream muggins.Stream) *muggins.Game {
	return &muggins.Game{
		NumPlayers: numPlayers,
		Players:    make([]muggins.Player, 0),
		Strm:       stream,
	}
}

func main() {
	gameDict := make(map[string]*muggins.Game)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		_, ok := gameDict["game"]
		if !ok {
			gameDict["game"] = NewGame(4, ws)
		}
		numClients := len(gameDict["game"].Players)
		if numClients == gameDict["game"].NumPlayers {
			return
		}
		clientID := r.Header.Get("sec-websocket-key")
		gameDict["game"].AddPlayer(muggins.Player{ID: clientID})
	})
	game := gameDict["game"]
	if game == nil {
		log.Print("Waiting for clients to connect.")
	} else {
		if game.NumPlayers < 2 || game.NumPlayers > 4 {
			log.Fatal("Too few or too many players")
		}
		game.Play()
		delete(gameDict, "game")
	}
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
