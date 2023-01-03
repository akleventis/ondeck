package main

import (
	"encoding/json"
	"net/http"
	"strconv"

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
		id := vars["id"]

		p, err := h.db.RetrievePerson(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if p == nil {
			apiResponse(w, http.StatusNotFound, lib.ErrPersonNotFound)
			return
		}

		apiResponse(w, 200, p)
	}
}

func (h *Handler) UpdatePerson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars == nil {
			http.Error(w, lib.ErrInvalidID.Error(), http.StatusBadRequest)
			return
		}
		id := vars["id"]

		p, err := h.db.RetrievePerson(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if p == nil {
			apiResponse(w, http.StatusNotFound, lib.ErrPersonNotFound)
			return
		}

		updatedPerson := &db.Person{
			ID: p.ID,
		}

		if err := json.NewDecoder(r.Body).Decode(&updatedPerson); err != nil {
			http.Error(w, lib.ErrInvalidArgJSONBody.Error(), http.StatusBadRequest)
			return
		}

		if updatedPerson.Name == "" {
			updatedPerson.Name = p.Name
		}
		if updatedPerson.Phone == 0 {
			updatedPerson.Phone = p.Phone
		}

		resPerson, err := h.db.UpdatePerson(updatedPerson)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		apiResponse(w, 200, resPerson)
	}
}

func (h *Handler) RemovePerson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars == nil {
			http.Error(w, lib.ErrInvalidID.Error(), http.StatusBadRequest)
			return
		}
		id := vars["id"]

		if err := h.db.RemovePerson(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		apiResponse(w, http.StatusGone, http.StatusText(410))
	}
}

func (h *Handler) RetrieveDrinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		drinks, err := h.db.RetrieveDrinks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		apiResponse(w, 200, drinks)
	}
}

func (h *Handler) CreateDrink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var d *db.Drink
		if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
			http.Error(w, lib.ErrInvalidArgJSONBody.Error(), http.StatusBadRequest)
			return
		}

		resDrink, err := h.db.CreateDrink(d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		apiResponse(w, 200, resDrink)
	}
}

func (h *Handler) RetrieveDrink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars == nil {
			http.Error(w, lib.ErrInvalidID.Error(), http.StatusBadRequest)
			return
		}
		id := vars["id"]

		resDrink, err := h.db.RetrieveDrink(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if resDrink == nil {
			apiResponse(w, http.StatusNotFound, lib.ErrDrinkNotFound)
			return
		}

		apiResponse(w, 200, resDrink)
	}
}

func (h *Handler) UpdateDrink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars == nil {
			http.Error(w, lib.ErrInvalidID.Error(), http.StatusBadRequest)
			return
		}
		id := vars["id"]

		d, err := h.db.RetrieveDrink(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if d == nil {
			apiResponse(w, http.StatusNotFound, lib.ErrDrinkNotFound)
			return
		}

		updatedDrink := &db.Drink{
			ID: d.ID,
		}

		if err := json.NewDecoder(r.Body).Decode(&updatedDrink); err != nil {
			http.Error(w, lib.ErrInvalidArgJSONBody.Error(), http.StatusBadRequest)
			return
		}

		if updatedDrink.Name == "" {
			updatedDrink.Name = d.Name
		}

		if updatedDrink.Price == 0 {
			updatedDrink.Price = d.Price
		}

		resDrink, err := h.db.UpdateDrink(updatedDrink)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		apiResponse(w, 200, resDrink)
	}
}

func (h *Handler) RemoveDrink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars == nil {
			http.Error(w, lib.ErrInvalidID.Error(), http.StatusBadRequest)
			return
		}
		id := vars["id"]

		if err := h.db.RemoveDrink(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		apiResponse(w, http.StatusGone, http.StatusText(410))
	}
}

func (h *Handler) CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars == nil {
			http.Error(w, lib.ErrInvalidID.Error(), http.StatusBadRequest)
			return
		}
		pID := vars["person_id"]

		p, err := h.db.RetrievePerson(pID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if p == nil {
			apiResponse(w, http.StatusNotFound, lib.ErrPersonNotFound)
			return
		}

		var o db.Order
		if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
			http.Error(w, lib.ErrInvalidArgJSONBody.Error(), http.StatusBadRequest)
			return
		}

		drinks := make([]db.DrinkOrder, 0)
		for _, v := range o.Order {
			id := strconv.Itoa(v.DrinkID)

			drink, err := h.db.RetrieveDrink(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if drink == nil {
				apiResponse(w, 404, lib.ErrDrinkNotFound)
				return
			}

			drinkOrder := &db.DrinkOrder{
				DrinkID: drink.ID,
				Name:    drink.Name,
				Price:   drink.Price,
				Comment: v.Comment,
			}
			drinks = append(drinks, *drinkOrder)
		}

		order := &db.FullOrder{
			Person: *p,
			Drinks: drinks,
		}

		resOrder, err := h.db.CreateOrder(order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		apiResponse(w, 200, resOrder)
	}
}

func (h *Handler) RemoveOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars == nil {
			http.Error(w, lib.ErrInvalidID.Error(), http.StatusBadRequest)
			return
		}
		orderNumber := vars["order_number"]

		err := h.db.RemoveOrder(orderNumber)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		apiResponse(w, http.StatusGone, http.StatusText(410))
	}
}

func (h *Handler) RetrieveOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars == nil {
			http.Error(w, lib.ErrInvalidID.Error(), http.StatusBadRequest)
			return
		}
		personID := vars["person_id"]

		orders, err := h.db.RetrieveOrders(personID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if orders == nil {
			apiResponse(w, http.StatusNotFound, http.StatusText(404))
		}

		apiResponse(w, 200, orders)
	}
}

func (h *Handler) RetrieveQueue() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
