package api

import "github.com/movieapi/usecase/movie"

type Module struct {
	usecase usecase
}

type usecase struct {
	movie movie.Usecase
}
