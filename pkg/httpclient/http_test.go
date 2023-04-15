package httpclient

import (
	"context"
	"net/http"
	"testing"
)

func TestDoRequestReal(t *testing.T) {
	ctx := context.Background()

	DoRequest(ctx, Request{
		ApiUrl: "http://www.omdbapi.com/",
		QueryParams: map[string]string{
			"apikey": "5a7ea26b",
			"t":      "raid",
		},
		Method: http.MethodGet,
	})
}
