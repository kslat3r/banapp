package games

import (
	"ban-app/common"

	"github.com/globalsign/mgo/bson"
)

type player struct {
	GameID  bson.ObjectId  `json:"gameId" bson:"gameId"`
	ID      bson.ObjectId  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string         `json:"name" binding:"required" bson:"name"`
	Letters map[string]int `json:"letters" binding:"required" bson:"letters"`
}

func createPlayer(gameID bson.ObjectId, p player) player {
	p.GameID = gameID
	p.ID = bson.NewObjectId()
	p.Letters = map[string]int{}

	err := common.Db.C("players").Insert(p)

	if err != nil {
		panic(err)
	}

	return p
}
