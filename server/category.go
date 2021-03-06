package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/muathendirangu/lavida-api/category"
)

//Handler handles incoming requests to category service/usecase
type categoryHandler struct {
	usecase category.Service
}

func (c *categoryHandler) router() chi.Router {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Post("/", c.addCategory)
		r.Get("/", c.getCategories)
		r.Get("/{ID}", c.getCategory)
	})
	return r
}

func (c *categoryHandler) addCategory(w http.ResponseWriter, r *http.Request) {
	var category category.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		log.Fatal(err)
	}
	err = c.usecase.AddCategory(&category)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte("category created"))

}

func (c *categoryHandler) getCategories(w http.ResponseWriter, r *http.Request) {
	var categories []*category.Category
	categories = c.usecase.GetCategories()
	payload, err := json.Marshal(categories)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(payload)
}

func (c *categoryHandler) getCategory(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "ID")
	catID, err := strconv.Atoi(ID)
	if err != nil {
		log.Fatal(err)
	}
	category := c.usecase.GetCategory(catID)
	payload, err := json.Marshal(category)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(payload)
}
