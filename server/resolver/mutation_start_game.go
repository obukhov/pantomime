package resolver

import "context"

type TeamInput struct {
	Name    string
	Members []string
}

type OptionsInput struct {
	RoundTime int32
}

type StartGameArgs struct {
	Teams   []*TeamInput
	Options *OptionsInput
}

func (r *Resolver) StartGame(ctx context.Context, args StartGameArgs) (*Game, error) {

	return &Game{}, nil
}
