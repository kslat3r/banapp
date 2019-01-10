package games

import (
	"banapp/common"

	"github.com/globalsign/mgo/bson"
)

func create(data *gameData) (*game, error) {
	game := game{}
	game.ID = bson.NewObjectId()
	game.Name = data.Name
	game.RemainingLetters = letters
	game.Status = statusWaiting
	game.Players = []player{}

	err := common.Db.C("games").Insert(&game)

	if err != nil {
		return nil, err
	}

	return &game, nil
}

func list() ([]game, error) {
	games := []game{}

	err := common.Db.C("games").Find(nil).All(&games)

	if err != nil {
		return nil, err
	}

	return games, nil
}

func getByID(id string) (*game, error) {
	game := game{}

	err := common.Db.C("games").FindId(bson.ObjectIdHex(id)).One(&game)

	if err != nil {
		return nil, err
	}

	return &game, nil
}
