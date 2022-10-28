package app

import "github.com/mux"

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()
}
