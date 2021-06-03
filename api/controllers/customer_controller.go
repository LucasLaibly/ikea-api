package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/LucasLaibly/ikea-api/api/models"
	"github.com/LucasLaibly/ikea-api/api/responses"
	"github.com/LucasLaibly/ikea-api/api/utils/formaterror"
	"github.com/gorilla/mux"
)

/*
Create customer record
*/
func (server *Server) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customer := models.Customer{}

	err = json.Unmarshal(body, &customer)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customer.Prepare()

	err = customer.Validate("")

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	customerCreated, err := customer.SaveCustomer(server.DB)

	if err != nil {
		_ = formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.RequestURI, customer.ID))
	responses.JSON(w, http.StatusCreated, customerCreated)
}

/*
Find customer by id
*/
func (server *Server) FindCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	customer := models.Customer{}

	customerFound, err := customer.FindCustomerByID(server.DB, string(uid))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, customerFound)
}
