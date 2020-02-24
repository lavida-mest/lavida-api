package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/muathendirangu/lavida-api/trip"
)

//Handler handles incoming requests to guide service/usecase
type tripHandler struct {
	usecase trip.Service
}

func (g *tripHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Post("/", g.addTrip)
	})
	return r
}

func (g *tripHandler) addTrip(w http.ResponseWriter, r *http.Request) {
	var trip trip.Trip
	err := json.NewDecoder(r.Body).Decode(&trip)
	if err != nil {
		log.Fatal(err)
	}
	err = g.usecase.AddTrip(&trip)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte("trip created"))

}
