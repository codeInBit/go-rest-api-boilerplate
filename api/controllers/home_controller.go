package controllers

import (
	"github.com/codeinbit/go-rest-api-boilerplate/api/utilities"
	"net/http"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request)  {
	utilities.JSON(w, http.StatusOK, "Welcome to Go Shop API")
}
