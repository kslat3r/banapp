package games

import (
	"ban-app/common"

	"github.com/globalsign/mgo/bson"
)

type game struct {
	ID   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" binding:"required" bson:"name"`
}

func list() []game {
	var list = []game{}

	err := common.Db.C("games").Find(nil).All(&list)

	if err != nil {
		panic(err)
	}

	return list
}

func create(g game) game {
	g.ID = bson.NewObjectId()

	err := common.Db.C("games").Insert(g)

	if err != nil {
		panic(err)
	}

	return g
}
