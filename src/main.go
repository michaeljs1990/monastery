package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/michaeljs1990/monastery/src/handler"
)

func main() {

	relic := mux.NewRouter()

	relic.HandleFunc("/upload/{bucket}/{name}", handler.DefaultUpload)
	relic.HandleFunc("/download/{bucket}/{name}", handler.DefaultDownload)

	http.ListenAndServe(":8080", relic)
}
