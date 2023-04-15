package httpclient

import "context"

type HttpClient interface {
	DoRequest(ctx context.Context, request Request) (map[string]interface{}, error)
}
