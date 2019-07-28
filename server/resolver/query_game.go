package resolver

import (
	"context"
	"github.com/graph-gophers/graphql-go"
)

type GameArgs struct {
	Id graphql.ID
}

func (r *Resolver) Game(ctx context.Context, args GameArgs) ([]*Game, error) {
	result := make([]*Game, 0)

	result = append(result, &Game{})

	return result, nil
}
