package route

import (
	"github.com/rs/cors"
	"net/http"
)

type Conf struct {
	Path string
	F func(w http.ResponseWriter, r *http.Request)
}

func Route(conf []Conf) http.Handler{
	r := http.NewServeMux()
	for _, route := range conf{
		r.HandleFunc(route.Path, route.F)
	}
	return cors.Default().Handler(r)
}