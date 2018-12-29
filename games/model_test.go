package games

import "testing"

func TestGetAll(t *testing.T) {
	games := getAll()

	if len(games) != len(all) {
		t.Fail()
	}

	for key, game := range games {
		if game.ID != all[key].ID {
			t.Fail()

			break
		}
	}
}
