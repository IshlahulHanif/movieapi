package api

import (
	"github.com/IshlahulHanif/poneglyph"
	"github.com/movieapi/usecase/movie"
	"github.com/movieapi/utils"
	"sync"
)

var (
	m    Module
	once sync.Once
)

func GetInstance(c utils.Config) (Module, error) {
	var (
		errFinal error
	)

	// TODO: in this case, if there is an init error and m is not initialized, we can never init this instance
	once.Do(func() {
		movieUc, err := movie.GetInstance(c)
		if err != nil {
			errFinal = poneglyph.Trace(err)
			return
		}

		m = Module{
			usecase: usecase{
				movie: movieUc,
			},
		}
	})

	return m, errFinal
}
