package main

import (
	"encoding/json"
	"net/http"

	"github.com/akleventis/ondeck/db"
	"github.com/akleventis/ondeck/lib"
	"github.com/gorilla/mux"
)

type Handler struct {
	db *db.DB
}

func NewHandler(db *db.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func apiResponse(w http.ResponseWriter, code int, obj interface{}) {
	r, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(r)
}

func (h *Handler) CreatePerson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p *db.Person
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, lib.ErrInvalidArgJSONBody.Error(), http.StatusBadRequest)
			return
		}

		p, err := h.db.CreatePerson(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		apiResponse(w, 200, p)
	}
}

func (h *Handler) RetrievePerson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars == nil {
			http.Error(w, lib.ErrInvalidID.Error(), http.StatusBadRequest)
			return
		}
		// id := vars["id"]
	}
}

func (h *Handler) UpdatePerson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars == nil {
			http.Error(w, lib.ErrInvalidID.Error(), http.StatusBadRequest)
			return
		}
		// id := vars["id"]
		// get person to update
		// decode body
		// update
	}
}

func (h *Handler) RemovePerson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars == nil {
			http.Error(w, lib.ErrInvalidID.Error(), http.StatusBadRequest)
			return
		}
		// id := vars["id"]
		// db delete by id
	}
}

func (h *Handler) RetrieveDrinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// return all drinks stored in db
	}
}

func (h *Handler) CreateDrink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var d db.Drink
		if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
			http.Error(w, lib.ErrInvalidArgJSONBody.Error(), http.StatusBadRequest)
			return
		}
		// create drink
	}
}

func (h *Handler) RetrieveDrink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars == nil {
			http.Error(w, lib.ErrInvalidID.Error(), http.StatusBadRequest)
			return
		}
		// drink := vars["name"]
		// retrieve drink by name
	}
}

func (h *Handler) UpdateDrink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars == nil {
			http.Error(w, lib.ErrInvalidID.Error(), http.StatusBadRequest)
			return
		}
		// drink := vars["name"]
		// get drink to update
		// decode body
		// update
	}
}

func (h *Handler) RemoveDrink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars == nil {
			http.Error(w, lib.ErrInvalidID.Error(), http.StatusBadRequest)
			return
		}
		// drink := vars["name"]
		// remove from db
	}
}

func (h *Handler) CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var o db.Order
		if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
			http.Error(w, lib.ErrInvalidArgJSONBody.Error(), http.StatusBadRequest)
			return
		}
		// retrieve person by id
		// create Order
		// retrieve serial "order number", append to Order obj and return
	}
}

func (h *Handler) RemoveOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars == nil {
			http.Error(w, lib.ErrInvalidID.Error(), http.StatusBadRequest)
			return
		}
		// order := vars["order_number"]
		// remove from db
	}
}

func (h *Handler) RetrieveQueue() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) RetrieveOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars == nil {
			http.Error(w, lib.ErrInvalidID.Error(), http.StatusBadRequest)
			return
		}
		// person := vars["person_id"]
		// retrieve all orders by person
	}
}
