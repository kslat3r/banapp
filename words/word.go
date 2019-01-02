package words

import (
	"fmt"
	"net/http"
	"os"
)

type word struct {
	Value  string `json:"value"`
	Exists bool   `json:"exists"`
}

func get(value string) word {
	appID := os.Getenv("OXFORD_APP_ID")
	appSecret := os.Getenv("OXFORD_APP_SECRET")

	if len(appID) == 0 || len(appSecret) == 0 {
		panic("Could not get Oxford dictionary app info")
	}

	req, reqErr := http.NewRequest("GET", fmt.Sprintf("https://od-api.oxforddictionaries.com:443/api/v1/entries/en/%s", value), nil)
	req.Header.Set("app_id", appID)
	req.Header.Set("app_key", appSecret)

	if reqErr != nil {
		panic(reqErr)
	}

	client := &http.Client{}
	res, _ := client.Do(req)

	exists := false

	if res.StatusCode != 404 {
		exists = true
	}

	return word{
		Value:  value,
		Exists: exists,
	}
}
