package httpclient

import (
	"context"
	"encoding/json"
	"github.com/IshlahulHanif/poneglyph"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (m Module) DoRequest(ctx context.Context, request Request) (map[string]interface{}, error) {
	var result = make(map[string]interface{})

	// create a new request
	apiURL := request.ApiUrl

	// insert query params
	if len(request.QueryParams) > 0 {
		queryParams := url.Values{}
		for key, param := range request.QueryParams {
			queryParams.Add(key, param)
		}
		apiURL += "?" + queryParams.Encode()
	}

	// Create an HTTP request object with the URL and method
	req, err := http.NewRequest(request.Method, apiURL, nil)
	if err != nil {
		return result, poneglyph.Trace(err)
	}
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return result, poneglyph.Trace(err)
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, poneglyph.Trace(err)
	}

	if err = json.Unmarshal(contents, &result); err != nil {
		return result, poneglyph.Trace(err)
	}

	return result, nil
}
