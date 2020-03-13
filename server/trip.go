package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
		r.Get("/search/{traveler_type}&{trip_month}&{trip_activity}", g.searchTrip)
		r.Get("/{ID}&{guideID}", g.viewTrip)
		r.Get("/", g.getTrips)
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
	Traveler := chi.URLParam(r, "traveler_type")
	Month := chi.URLParam(r, "trip_month")
	Activity := chi.URLParam(r, "trip_activity")
	trips := g.usecase.SearchTrip(Traveler, Month, Activity)
	payload, err := json.Marshal(trips)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(payload)
}

func (g *tripHandler) viewTrip(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "ID")
	guide := chi.URLParam(r, "guideID")
	tripID, err := strconv.Atoi(ID)
	if err != nil {
		log.Fatal(err)
	}
	guideID, err := strconv.Atoi(guide)
	if err != nil {
		log.Fatal(err)
	}

	trips := g.usecase.ViewTrip(int64(tripID), int64(guideID))
	payload, err := json.Marshal(trips)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(payload)
}

func (g *tripHandler) getTrips(w http.ResponseWriter, r *http.Request) {
	var trips []*trip.Trip
	trips = g.usecase.GetTrips()
	payload, err := json.Marshal(trips)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(payload)
}
