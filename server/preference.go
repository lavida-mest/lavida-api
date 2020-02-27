package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/muathendirangu/lavida-api/preference"
)

//Handler handles incoming requests to preference service/usecase
type preferenceHandler struct {
	usecase preference.Service
}

func (p *preferenceHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Post("/", p.savePreference)
	})
	return r
}

func (p *preferenceHandler) savePreference(w http.ResponseWriter, r *http.Request) {
	var preference preference.Preference
	err := json.NewDecoder(r.Body).Decode(&preference)
	if err != nil {
		log.Fatalf("An error occurred while decoding the request to preference struct: %v", err)
	}
	err = p.usecase.AddPreference(&preference)
	if err != nil {
		log.Fatalf("An error occurred while handling the save preference usecase: %v", err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte("preference saved"))
}
