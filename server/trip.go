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
		r.Get("/search/{trip_location}&{trip_duration}&{traveler_type}&{trip_month}&{trip_year}", g.searchTrip)
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

func (g *tripHandler) searchTrip(w http.ResponseWriter, r *http.Request) {
	Location := chi.URLParam(r, "trip_location")
	Duration := chi.URLParam(r, "trip_duration")
	Traveler := chi.URLParam(r, "traveler_type")
	Month := chi.URLParam(r, "trip_month")
	Year := chi.URLParam(r, "trip_year")
	trips := g.usecase.SearchTrip(Location, Duration, Traveler, Month, Year)
	payload, err := json.Marshal(trips)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(payload)
}
