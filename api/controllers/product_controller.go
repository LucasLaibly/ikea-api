package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/LucasLaibly/ikea-api/api/models"
	"github.com/LucasLaibly/ikea-api/api/responses"
	"github.com/LucasLaibly/ikea-api/api/utils/formaterror"
	"github.com/gorilla/mux"
)

/*
Create Product
*/
func (server *Server) CreateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	product := models.Product{}

	err = json.Unmarshal(body, &product)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	product.Prepare()

	err = product.Validate("")

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	productCreated, err := product.SaveProduct(server.DB)

	if err != nil {
		foramtedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, foramtedError)
	}

	w.Header().Set("Location: ", fmt.Sprintf("%s%s/%s", r.Host, r.URL.Path, productCreated.ID))
	responses.JSON(w, http.StatusCreated, productCreated)
}

/*
Find Product
*/
func (server *Server) FindProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	product := models.Product{}

	toFind, err := product.FindProductByID(server.DB, id)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, toFind)
}
