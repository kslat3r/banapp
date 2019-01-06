package games

import (
	"math/rand"
	"reflect"
	"time"
)

func dealTiles(game *game) {
	rand.Seed(time.Now().Unix())

	// get number of players

	numPlayers := len(game.Players)

	// get num of tiles

	var numTiles int

	if numPlayers >= 2 && numPlayers <= 4 {
		numTiles = 21
	} else if numPlayers > 4 && numPlayers <= 6 {
		numTiles = 15
	} else if numPlayers > 7 {
		numTiles = 11
	}

	for _, player := range game.Players {
		for i := 0; i < numTiles; i++ {
			// get random index

			keys := reflect.ValueOf(game.RemainingLetters).MapKeys()
			random := rand.Int() % len(keys)

			// get letter

			letter := keys[random].String()

			// add to player

			if _, ok := player.Letters[letter]; !ok {
				player.Letters[letter] = 0
			}

			// counts

			player.Letters[letter]++
			game.RemainingLetters[letter]--

			// if remaining is zero, remove from map

			if game.RemainingLetters[letter] == 0 {
				delete(game.RemainingLetters, letter)
			}
		}
	}
}
