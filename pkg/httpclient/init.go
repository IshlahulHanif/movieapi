package httpclient

import (
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
		m = Module{}
	})

	return m, errFinal
}
