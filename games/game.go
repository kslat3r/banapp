package games

import (
	"ban-app/common"
	"fmt"

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

type game struct {
	ID               bson.ObjectId  `json:"_id" binding:"required" bson:"_id"`
	Name             string         `json:"name" binding:"required" bson:"name"`
	RemainingLetters map[string]int `json:"remainingLetters" binding:"required" bson:"remainingLetters"`
	Status           status         `json:"status" binding:"required" bson:"status"`
}

func createGame(g game) game {
	g.ID = bson.NewObjectId()
	g.RemainingLetters = letters
	g.Status = statusWaiting

	err := common.Db.C("games").Insert(g)

	if err != nil {
		panic(err)
	}

	return g
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

	fmt.Println(err)

	if err != nil {
		panic(err)
	}

	return game
}

func updateGame(id bson.ObjectId, g game) game {
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": g.Status}}
	err := common.Db.C("games").Update(query, update)

	if err != nil {
		panic(err)
	}

	return getGame(id.Hex())
}
