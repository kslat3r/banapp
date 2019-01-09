package common

import (
	"fmt"
	"os"

	"github.com/globalsign/mgo"
)

const defaultURL = "mongodb://localhost:27017"
const defaultDb = "banapp"

// Session the global session
var Session *mgo.Session

// Db the DB instance
var Db *mgo.Database

// InitDb initialise everything
func InitDb() {
	url := fmt.Sprintf("%s/%s", defaultURL, defaultDb)

	user := os.Getenv("BANAPP_MONGODB_USER")
	password := os.Getenv("BANAPP_MONGODB_PASSWORD")
	hostname := os.Getenv("BANAPP_MONGODB_SERVICE_HOST")
	port := os.Getenv("BANAPP_MONGODB_SERVICE_PORT")

	if len(user) > 0 && len(password) > 0 && len(hostname) > 0 && len(port) > 0 {
		url = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", user, password, hostname, port, defaultDb)
	}

	info, err := mgo.ParseURL(url)

	if err != nil {
		panic(err)
	}

	Session, err := mgo.DialWithInfo(info)

	if err != nil {
		panic(err)
	}

	Session.SetSafe(&mgo.Safe{})

	Db = Session.DB("banapp")
}
