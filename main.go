package main

import (
	"log"
	"net/http"

	"github.com/akleventis/ondeck/db"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	db     *db.DB
}

func main() {
	port := ":8080"

	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := &server{
		router: mux.NewRouter(),
		db:     db,
	}

	s.router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}).Methods("GET")

	h := NewHandler(s.db)
	s.router.HandleFunc("/person", h.CreatePerson()).Methods("POST")
	s.router.HandleFunc("/person/{id}", h.RetrievePerson()).Methods("GET")
	s.router.HandleFunc("/person/{id}", h.UpdatePerson()).Methods("PATCH")
	s.router.HandleFunc("/person/{id}", h.RemovePerson()).Methods("DELETE")

	s.router.HandleFunc("/drinks", h.RetrieveDrinks()).Methods("GET")
	s.router.HandleFunc("/drink", h.CreateDrink()).Methods("POST")
	s.router.HandleFunc("/drink/{id}", h.RetrieveDrink()).Methods("GET")
	s.router.HandleFunc("/drink/{id}", h.UpdateDrink()).Methods("PATCH")
	s.router.HandleFunc("/drink/{id}", h.RemoveDrink()).Methods("DELETE")

	s.router.HandleFunc("/order/{person_id}", h.CreateOrder()).Methods("POST")
	s.router.HandleFunc("/order/{order_number}", h.RemoveOrder()).Methods("DELETE")

	s.router.HandleFunc("/orders/{person_id}", h.RetrieveOrders()).Methods("GET")

	s.router.HandleFunc("/queue", h.RetrieveQueue()).Methods("GET")

	handler := cors.Default().Handler(s.router)

	logrus.Infof("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
