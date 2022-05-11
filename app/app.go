package app

import (
	"github.com/comfysweet/bookstore_items-api/clients/elastic_search"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	Router = mux.NewRouter()
)

func StartApplication() {
	elastic_search.Init()
	mapUrls()

	srv := &http.Server{
		Handler: Router,
		Addr:    "127.0.0.1:8000",
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
