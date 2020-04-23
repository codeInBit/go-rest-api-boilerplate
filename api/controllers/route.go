package controllers

import (
	"github.com/codeinbit/go-rest-api-boilerplate/api/middlewares"
)

func (s *Server) LoadRoutes() {
	//Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
}