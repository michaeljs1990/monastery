package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/michaeljs1990/monastery/src/handler"

	_ "github.com/michaeljs1990/monastery/src/config"
)

func main() {
	// Handle parsing of flags that are loaded in via the init function
	// inside files for the config directory.
	flag.Parse()

	relic := mux.NewRouter()

	relic.HandleFunc("/upload/{bucket}/{name}", handler.DefaultUpload)
	relic.HandleFunc("/download/{bucket}/{name}", handler.DefaultDownload)

	http.ListenAndServe(":8080", relic)
}
