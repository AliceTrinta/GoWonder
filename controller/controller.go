package controller

import (
	"GoWonder/handler"
	"GoWonder/route"
	"net/http"
)

//Here, at the Begin Function, the routes are configured.
func Begin() http.Handler {

	routes := []route.Conf{
		{Path: "/", F: handler.Index, Method: "GET"},
		{Path: "/create", F: handler.Create, Method: "POST"},
		{Path: "/all", F: handler.GetAll, Method: "GET"},
		{Path: "/get/{id}", F: handler.GetByID, Method: "GET"},
		{Path: "/up/{id}", F: handler.Update, Method: "PUT"},
		{Path: "/del/{id}", F: handler.Delete, Method: "DELETE"},
	}

	return route.Route(routes)
}
