package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/michaeljs1990/monastery/src/config"
	"github.com/michaeljs1990/monastery/src/handler"
)

func main() {
	// Handle parsing of flags that are loaded in via the init function
	// inside files for the config directory.
	flag.Parse()

	relic := mux.NewRouter()

	relic.HandleFunc("/upload/{bucket}/{name}", handler.DefaultUpload)
	relic.HandleFunc("/download/{name}", handler.DefaultDownload)

	http.ListenAndServe(":"+config.ServicePort, relic)
}
