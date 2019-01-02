package words

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type dictionary struct {
	Results []struct {
		Word string `json:"word"`
	}
}

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

	req, reqErr := http.NewRequest("GET", fmt.Sprintf("https://od-api.oxforddictionaries.com:443/api/v1/search/en?q=%s&prefix=false", value), nil)
	req.Header.Set("app_id", appID)
	req.Header.Set("app_key", appSecret)

	if reqErr != nil {
		panic(reqErr)
	}

	client := &http.Client{}
	res, resErr := client.Do(req)

	if resErr != nil {
		panic(resErr)
	}

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		panic(readErr)
	}

	dict := dictionary{}
	jsonErr := json.Unmarshal(body, &dict)

	if jsonErr != nil {
		panic(jsonErr)
	}

	exists := false

	if len(dict.Results) > 0 {
		exists = true
	}

	return word{
		Value:  value,
		Exists: exists,
	}
}
