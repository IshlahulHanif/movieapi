package movie

import (
	"github.com/IshlahulHanif/poneglyph"
	"github.com/movieapi/pkg/httpclient"
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

	once.Do(func() {
		http, err := httpclient.GetInstance(c)
		if err != nil {
			errFinal = poneglyph.Trace(err)
			return
		}

		m = Module{
			endpoint: endpoint{
				http: http,
			},
		}
	})

	return m, errFinal
}
