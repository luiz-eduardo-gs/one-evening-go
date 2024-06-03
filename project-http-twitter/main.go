package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"twitter/server"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	s := server.Server{
		Repository: &server.TweetMemoryRepository{},
	}

	go spamTweets()

	r.Get("/tweets", s.ListTweets)
	r.With(httprate.LimitByIP(10, 1*time.Minute)).Post("/tweets", s.AddTweet)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func spamTweets() {
	url := "http://localhost:8080/tweets"

	tw := server.Tweet{
		Message:  "ass",
		Location: "ass",
	}

	payload, err := json.Marshal(tw)
	if err != nil {
		log.Println("Error marshal")
		return
	}

	for {
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
		if err != nil {
			return
		}
		defer resp.Body.Close()
	}
}
