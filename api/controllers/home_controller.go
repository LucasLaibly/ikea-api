package controllers

import (
	"net/http"

	"github.com/LucasLaibly/ikea-api/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to IKEA, now serving you immediately.")
}
