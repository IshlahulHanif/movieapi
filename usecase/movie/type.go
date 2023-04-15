package movie

import "github.com/movieapi/pkg/httpclient"

type Module struct {
	endpoint endpoint
}

type endpoint struct {
	http httpclient.HttpClient
}
