package games

import (
	"banapp/common"
	"errors"

	"github.com/globalsign/mgo/bson"
)

type game struct {
	ID               bson.ObjectId  `json:"_id" bson:"_id"`
	Name             string         `json:"name" bson:"name"`
	RemainingLetters map[string]int `json:"remainingLetters" bson:"remainingLetters"`
	Status           status         `json:"status" bson:"status"`
	Players          []player       `json:"players" bson:"players"`
}

var letters = map[string]int{
	"A": 13,
	"B": 3,
	"C": 3,
	"D": 6,
	"E": 18,
	"F": 3,
	"G": 4,
	"H": 3,
	"I": 12,
	"J": 2,
	"K": 2,
	"L": 5,
	"M": 3,
	"N": 8,
	"O": 11,
	"P": 3,
	"Q": 2,
	"R": 9,
	"S": 6,
	"T": 9,
	"U": 6,
	"V": 3,
	"W": 3,
	"X": 2,
	"Y": 3,
	"Z": 2,
}

type status int

const (
	statusWaiting   status = 0
	statusPlaying   status = 1
	statusCompleted status = 2
)

type player struct {
	Name    string         `json:"name" bson:"name"`
	Letters map[string]int `json:"letters" bson:"letters"`
}

func (game *game) update(data *gameData) (*game, error) {
	if game.Status == statusPlaying && data.Status == statusWaiting {
		return nil, errors.New("Game is not in correct status to change to waiting")
	}

	if game.Status == statusWaiting && data.Status == statusPlaying {
		if len(game.Players) >= 2 {
			dealTiles(game)
		} else {
			return nil, errors.New("Game does not have enough players to begin")
		}
	}

	game.Status = data.Status

	err := common.Db.C("games").Update(bson.M{"_id": game.ID}, bson.M{"$set": game})

	if err != nil {
		return nil, err
	}

	return game, nil
}

func (game *game) addPlayer(data *playerData) (*game, error) {
	if game.Status != statusWaiting {
		return nil, errors.New("Game is not in correct status to add player")
	}

	player := player{Name: data.Name, Letters: map[string]int{}}
	game.Players = append(game.Players, player)

	err := common.Db.C("games").Update(bson.M{"_id": game.ID}, bson.M{"$set": game})

	if err != nil {
		return nil, err
	}

	return game, nil
}
