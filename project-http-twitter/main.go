package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	s := server{
		repository: &tweetMemoryRepository{},
	}

	http.HandleFunc("/tweets", s.tweets)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type response struct {
	ID int `json:"ID"`
}

type listTweetsResponse struct {
	Tweets []tweet `json:"tweets"`
}

type tweetRepository interface {
	Add(u tweet) (int, error)
	Tweets() ([]tweet, error)
}

type tweetMemoryRepository struct {
	tweets []tweet
}

func (r *tweetMemoryRepository) Add(tw tweet) (int, error) {
	r.tweets = append(r.tweets, tw)
	return len(r.tweets), nil
}

func (r *tweetMemoryRepository) Tweets() ([]tweet, error) {
	return r.tweets, nil
}

type server struct {
	repository tweetRepository
}

func (s server) addTweet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read body: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	tw := tweet{}

	if err := json.Unmarshal(body, &tw); err != nil {
		log.Println("Failed to unmarshal payload: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if tw.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("Tweet: `%s` from %s\n", tw.Message, tw.Location)

	id, err := s.repository.Add(tw)
	if err != nil {
		log.Println("Failed to create tweet: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := response{ID: id}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.Println("Failed to marshal: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id++
	w.Write(respJSON)
}

func (s server) listTweets(w http.ResponseWriter, r *http.Request) {
	tweets, err := s.repository.Tweets()

	resp := listTweetsResponse{
		Tweets: tweets,
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.Println("Failed to marshal tweets: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(respJSON)
}

func (s server) tweets(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.addTweet(w, r)
	} else if r.Method == http.MethodGet {
		s.listTweets(w, r)
	}
}
