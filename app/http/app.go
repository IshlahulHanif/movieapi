package main

import (
	"fmt"
	"github.com/IshlahulHanif/poneglyph"
	"github.com/movieapi/api"
	"github.com/movieapi/utils"
	"net/http"
)

func main() {
	// init config
	conf := utils.Config{}

	// init poneglyph settings
	poneglyph.SetProjectName("movieapi")
	poneglyph.SetIsPrintFromContentRoot(true)
	poneglyph.SetIsPrintFunctionName(true)
	poneglyph.SetIsPrintNewline(true)
	poneglyph.SetIsUseTabSeparator(false)

	var httpApi api.API
	// init http api
	httpApi, err := api.GetInstance(conf)
	if err != nil {
		err = poneglyph.Trace(err)
		fmt.Println(poneglyph.GetErrorLogMessage(err))
		return
	}

	// register handlers
	http.HandleFunc("/search", httpApi.HandlerSearch)
	http.HandleFunc("/detail/", httpApi.HandlerDetail)

	err = http.ListenAndServe("", nil)
	if err != nil {
		err = poneglyph.Trace(err)
		fmt.Println(poneglyph.GetErrorLogMessage(err))
		return
	}
}
