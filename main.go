package main

import (
	"fmt"
	"go-shorten-link/handler"
	"go-shorten-link/repository"
	"go-shorten-link/service"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"github.com/rs/cors"
)

func main() {

	// rd client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// dependency inject

	shortenRepository := repository.NewRedisRepository(rdb)
	shortenService := service.NewRedisService(shortenRepository)
	shortenHandler := handler.NewRedisShortenLinkHandler(shortenService)

	r := mux.NewRouter()

	r.HandleFunc("/shorten-link", shortenHandler.CreateShortenLink).Methods("POST")
	r.HandleFunc("/list-key/list", shortenHandler.ListLink).Methods("GET")
	r.HandleFunc("/{key}", shortenHandler.ResolveShortenLink).Methods("GET")

	// r.Use(CORSMiddleware())
	fmt.Print("this app running on :8080")
	handler := cors.Default().Handler(r)

	http.ListenAndServe(":9000", handler)
}
