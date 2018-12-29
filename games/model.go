package games

type game struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var all = []game{
	game{
		ID:   1,
		Name: "One",
	},
	game{
		ID:   2,
		Name: "Two",
	},
}

func getAll() []game {
	return all
}

func create(g game) {
	all = append(all, g)
}
