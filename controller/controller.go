package controller

import (
	"GoWonder/handler"
	"GoWonder/route"
	"net/http"
)

func Begin() http.Handler{

	routes := []route.Conf{
		{"/", handler.Index},
	}

	return route.Route(routes)
}

