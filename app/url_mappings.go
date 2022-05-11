package app

import (
	"github.com/comfysweet/bookstore_items-api/controllers"
	"net/http"
)

func mapUrls() {
	Router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	Router.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)
	Router.HandleFunc("/items/{id}", controllers.ItemsController.Search).Methods(http.MethodGet)
}
