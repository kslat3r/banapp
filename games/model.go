package games

import "github.com/rs/xid"

type game struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var all = []game{}

func list() []game {
	return all
}

func create(g game) game {
	g.ID = xid.New().String()

	all = append(all, g)

	return g
}
