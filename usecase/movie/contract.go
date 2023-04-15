package movie

import "context"

type Usecase interface {
	SearchByName(ctx context.Context, movieName string) (map[string]interface{}, error)
	GetDetailByID(ctx context.Context, id string) (map[string]interface{}, error)
}
