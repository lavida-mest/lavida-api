package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/muathendirangu/lavida-api/category"
)

//Server holds dependencies for http server
type Server struct {
	Category category.Service
	router   chi.Router
}

//New returns a new HTTP server
func New(cs category.Service) *Server {
	s := &Server{
		Category: cs,
	}
	r := chi.NewRouter()
	r.Use(accessControl)
	r.Route("/trip", func(r chi.Router) {
		h := categoryHandler{s.Category}
		r.Mount("/", h.router())
	})
	s.router = r
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
