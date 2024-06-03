package server

import "sync"

type tweetRepository interface {
	Add(u Tweet) (int, error)
	Tweets() ([]Tweet, error)
}

type TweetMemoryRepository struct {
	tweets []Tweet
	lock   sync.RWMutex
}

func (r *TweetMemoryRepository) Add(tw Tweet) (int, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.tweets = append(r.tweets, tw)
	return len(r.tweets), nil
}

func (r *TweetMemoryRepository) Tweets() ([]Tweet, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	return r.tweets, nil
}
