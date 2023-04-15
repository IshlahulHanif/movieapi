package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/IshlahulHanif/poneglyph"
	"github.com/movieapi/entity"
	"net/http"
	"strings"
	"time"
)

func (m Module) HandlerSearch(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  context.Context
		err  error
		resp entity.JSend
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // TODO: make const timeout
	defer cancel()

	if r.Method != http.MethodGet {
		err = poneglyph.Trace(errors.New("unsupported HTTP method"))
		fmt.Println(poneglyph.GetLogErrorTrace(err))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: err.Error(),
			},
		}

		if err = json.NewEncoder(w).Encode(resp); err != nil {
			err = poneglyph.Trace(err)
			fmt.Println(poneglyph.GetLogErrorTrace(err))
		}
		return
	}

	title := r.URL.Query().Get("title")

	result, err := m.usecase.movie.SearchByName(ctx, title)
	if err != nil { //TODO: standardize err response
		err = poneglyph.Trace(err)
		fmt.Println(poneglyph.GetLogErrorTrace(err))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: err.Error(),
			},
		}

		if err = json.NewEncoder(w).Encode(resp); err != nil {
			err = poneglyph.Trace(err)
			fmt.Println(poneglyph.GetLogErrorTrace(err))
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp = entity.JSend{
		Status: entity.ConstJSendSuccess,
		Data: struct {
			Result map[string]interface{} `json:"result"`
		}{
			Result: result,
		},
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		err = poneglyph.Trace(err)
		fmt.Println(poneglyph.GetLogErrorTrace(err))
	}
}

func (m Module) HandlerDetail(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  context.Context
		err  error
		resp entity.JSend
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // TODO: make const timeout
	defer cancel()

	if r.Method != http.MethodGet {
		err = poneglyph.Trace(errors.New("unsupported HTTP method"))
		fmt.Println(poneglyph.GetLogErrorTrace(err))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: err.Error(),
			},
		}

		if err = json.NewEncoder(w).Encode(resp); err != nil {
			err = poneglyph.Trace(err)
			fmt.Println(poneglyph.GetLogErrorTrace(err))
		}
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/detail/")

	result, err := m.usecase.movie.GetDetailByID(ctx, id)
	if err != nil { //TODO: standardize err response
		err = poneglyph.Trace(err)
		fmt.Println(poneglyph.GetLogErrorTrace(err))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp = entity.JSend{
			Status: entity.ConstJSendFail,
			Data: struct {
				ErrorMsg string `json:"error"`
			}{
				ErrorMsg: err.Error(),
			},
		}

		if err = json.NewEncoder(w).Encode(resp); err != nil {
			err = poneglyph.Trace(err)
			fmt.Println(poneglyph.GetLogErrorTrace(err))
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp = entity.JSend{
		Status: entity.ConstJSendSuccess,
		Data: struct {
			Result map[string]interface{} `json:"result"`
		}{
			Result: result,
		},
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		err = poneglyph.Trace(err)
		fmt.Println(poneglyph.GetLogErrorTrace(err))
	}
}
