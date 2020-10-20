package controller

import (
	"GoWonder/handler"
	"GoWonder/route"
	"net/http"
)

func Begin() http.Handler{

	routes := []route.Conf{
		{"/", handler.Index, "GET"},
		{"/create", handler.Create, "POST"},
		{"/all", handler.GetAll, "GET"},
		{"/get/{id}", handler.GetByID, "GET"},
		{"/up/{id}", handler.Update, "PUT"},
		{"/del/{id}", handler.Delete, "DELETE"},
	}

	return route.Route(routes)
}

