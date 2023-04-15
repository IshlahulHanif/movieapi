package movie

import (
	"context"
	"errors"
	"github.com/IshlahulHanif/poneglyph"
	"github.com/movieapi/pkg/httpclient"
	"net/http"
)

func (m Module) SearchByName(ctx context.Context, movieName string) (map[string]interface{}, error) {
	var (
		err  error
		resp = make(map[string]interface{})
	)

	if len(movieName) == 0 {
		err = poneglyph.Trace(errors.New("empty data"))
		return resp, err
	}

	resp, err = m.endpoint.http.DoRequest(ctx, httpclient.Request{
		ApiUrl: ConstOmdbApiUrl,
		QueryParams: map[string]string{
			"apikey": ConstApiKey,
			"t":      movieName,
		},
		Method: http.MethodGet,
	})
	if err != nil {
		return resp, poneglyph.Trace(err)
	}

	return resp, nil
}

func (m Module) GetDetailByID(ctx context.Context, id string) (map[string]interface{}, error) {
	var (
		err  error
		resp = make(map[string]interface{})
	)

	if len(id) == 0 {
		err = poneglyph.Trace(errors.New("empty data"))
		return resp, err
	}

	resp, err = m.endpoint.http.DoRequest(ctx, httpclient.Request{
		ApiUrl: ConstOmdbApiUrl,
		QueryParams: map[string]string{
			"apikey": ConstApiKey,
			"i":      id,
		},
		Method: http.MethodGet,
	})
	if err != nil {
		return resp, poneglyph.Trace(err)
	}

	return resp, nil
}
