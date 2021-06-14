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
Create Cart
*/
func (server *Server) CreateCart(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	//vars := mux.Vars(r)

	//id := vars["id"]

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	cart := models.Cart{}

	err = json.Unmarshal(body, &cart)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	cart.Prepare()

	cartCreated, err := cart.SaveCart(server.DB)

	if err != nil {
		formatedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formatedError)
	}

	w.Header().Set("Location: ", fmt.Sprintf("%s%s/%s", r.Host, r.URL.Path, cartCreated.ID))
	responses.JSON(w, http.StatusCreated, cartCreated)
}

/*
Get all cart items for a customer
*/
func (server *Server) GetAllCartItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	cart := models.Cart{}

	err := cart.GetAllItemsInCart(server.DB, id)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, cart)
}
