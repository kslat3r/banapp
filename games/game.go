package games

import (
	"bananapp/common"

	"github.com/globalsign/mgo/bson"
)

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

type game struct {
	ID               bson.ObjectId  `json:"_id" bson:"_id"`
	Name             string         `json:"name" bson:"name"`
	RemainingLetters map[string]int `json:"remainingLetters" bson:"remainingLetters"`
	Status           status         `json:"status" bson:"status"`
	Players          []player       `json:"players" bson:"players"`
}

func createGame(g *game) game {
	g.ID = bson.NewObjectId()
	g.RemainingLetters = letters
	g.Status = statusWaiting
	g.Players = []player{}

	err := common.Db.C("games").Insert(g)

	if err != nil {
		panic(err)
	}

	return *g
}

func getGames() []game {
	games := []game{}
	err := common.Db.C("games").Find(nil).All(&games)

	if err != nil {
		panic(err)
	}

	return games
}

func getGame(id string) game {
	game := game{}
	err := common.Db.C("games").FindId(bson.ObjectIdHex(id)).One(&game)

	if err != nil {
		panic(err)
	}

	return game
}

func updateGame(og *game, ng *game) game {
	if og.Status == statusWaiting && ng.Status == statusPlaying {
		for i := range og.Players {
			setStartingLetters(&og.Players[i], &og.RemainingLetters)
		}
	}

	query := bson.M{"_id": og.ID}
	update := bson.M{"$set": bson.M{"status": ng.Status, "remainingLetters": og.RemainingLetters, "players": og.Players}}
	err := common.Db.C("games").Update(query, update)

	if err != nil {
		panic(err)
	}

	return getGame(og.ID.Hex())
}

func addPlayer(g *game, p *player) game {
	if g.Status != statusWaiting {
		panic("Game is not in correct status to add player")
	}

	p.Letters = map[string]int{}

	query := bson.M{"_id": g.ID}
	update := bson.M{"$push": bson.M{"players": p}}
	err := common.Db.C("games").Update(query, update)

	if err != nil {
		panic(err)
	}

	return getGame(g.ID.Hex())
}

func setStartingLetters(p *player, letters *map[string]int) {
	p.Letters = *letters
}
