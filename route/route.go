package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Here, all the routes are aligned by the conf struct and configured by the rout function.

type Conf struct {
	Path   string
	F      func(w http.ResponseWriter, r *http.Request)
	Method string
}

func Route(conf []Conf) http.Handler {
	r := mux.NewRouter()
	for _, route := range conf {
		r.HandleFunc(route.Path, route.F).Methods(route.Method)
	}
	return cors.Default().Handler(r)
}
