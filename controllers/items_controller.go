package controllers

import (
	"encoding/json"
	"github.com/comfysweet/bookstore-oauth-go/oauth"
	"github.com/comfysweet/bookstore_items-api/domain/items"
	"github.com/comfysweet/bookstore_items-api/domain/queries"
	"github.com/comfysweet/bookstore_items-api/http_utils"
	"github.com/comfysweet/bookstore_items-api/services"
	"github.com/comfysweet/bookstore_utils-go/errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, errors.NewRestError(err.Message, err.Status, err.Error, nil))
		return
	}

	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		respErr := errors.NewNotFoundError("not found seller id")
		http_utils.RespondError(w, respErr)
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = sellerId

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, err := services.ItemsService.Get(itemId)
	if err != nil {
		http_utils.RespondError(w, err)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, item)
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, apiErr)
		return
	}
	defer r.Body.Close()

	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, apiErr)
	}

	items, searchErr := services.ItemsService.Search(query)
	if searchErr != nil {
		http_utils.RespondError(w, searchErr)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, items)
}
