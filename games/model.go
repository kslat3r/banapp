package games

type game struct {
	id int `json:"id"`
}

var all = []game{
	game{id: 1},
	game{id: 2},
}

func getAll() []game {
	return all
}
