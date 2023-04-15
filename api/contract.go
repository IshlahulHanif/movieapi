package api

import "net/http"

type API interface {
	HandlerSearch(w http.ResponseWriter, r *http.Request)
	HandlerDetail(w http.ResponseWriter, r *http.Request)
}
