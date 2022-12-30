package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	// db
}

func main() {
	port := ":8080"

	s := &server{
		// db:
		router: mux.NewRouter(),
	}

	s.router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	handler := cors.Default().Handler(s.router)

	logrus.Infof("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, handler))

}
