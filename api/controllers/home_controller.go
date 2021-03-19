package controllers

import (
	"net/http"

	"github.com/sang0968991331/service-golang/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
