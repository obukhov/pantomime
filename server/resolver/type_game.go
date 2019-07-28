package resolver

import (
	"github.com/graph-gophers/graphql-go"
)

type Game struct {
}

func (g *Game) Id() graphql.ID {
	return graphql.ID("1");
}
