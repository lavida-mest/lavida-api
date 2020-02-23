package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/muathendirangu/lavida-api/guide"
)

//Handler handles incoming requests to guide service/usecase
type guideHandler struct {
	usecase guide.Service
}

func (g *guideHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Post("/", g.addGuide)

	})
	return r
}

func (g *guideHandler) addGuide(w http.ResponseWriter, r *http.Request) {
	var guide guide.Guide
	err := json.NewDecoder(r.Body).Decode(&guide)
	if err != nil {
		log.Fatal(err)
	}
	err = g.usecase.AddGuide(&guide)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte("guide created"))

}
