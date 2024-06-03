package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Server struct {
	Repository tweetRepository
}

type Tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type response struct {
	ID int `json:"ID"`
}

type listTweetsResponse struct {
	Tweets []Tweet `json:"tweets"`
}

func (s Server) AddTweet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read body: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	tw := Tweet{}

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

	id, err := s.Repository.Add(tw)
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

func (s Server) ListTweets(w http.ResponseWriter, _ *http.Request) {
	tweets, _ := s.Repository.Tweets()

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
