package resolver

import "context"

type GamesArgs struct {
	Limit int32
	After *string
}

func (r *Resolver) Games(ctx context.Context, args GamesArgs) ([]*Game, error) {
	result := make([]*Game, 0)

	result = append(result, &Game{})

	return result, nil
}
