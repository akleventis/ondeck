package main

import (
	"log"
	"net/http"

	"github.com/akleventis/ondeck/db"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	db     *db.DB
}

func main() {
	port := ":8080"

	if err := godotenv.Load("./env"); err != nil {
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
	})

	handler := cors.Default().Handler(s.router)

	logrus.Infof("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
