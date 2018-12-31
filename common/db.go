package common

import (
	"os"

	"github.com/globalsign/mgo"
)

const mongoDbURL = "mongodb://localhost"

// Session the global session
var Session *mgo.Session

// Db the DB instance
var Db *mgo.Database

// InitDb initialise everything
func InitDb() {
	url := os.Getenv("MONGODB_URL")

	if len(url) == 0 {
		url = mongoDbURL
	}

	Session, err := mgo.Dial(url)

	if err != nil {
		panic(err)
	}

	Session.SetSafe(&mgo.Safe{})

	Db = Session.DB("bananagrams")
}
