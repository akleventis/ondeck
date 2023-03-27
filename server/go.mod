module github.com/akleventis/ondeck/server

go 1.18

require github.com/sirupsen/logrus v1.9.0

require (
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.4.0
	github.com/lib/pq v1.10.7
	github.com/rs/cors v1.8.3
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)

replace github.com/akleventis/ondeck/server => ./server
