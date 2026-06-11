package watcher

import (
	"encoding/json"
	"os"
	"github.com/google/uuid"
)

type Movie struct {
	Title       string   `json:"title"`
	Genres      []string `json:"genres"`
	ReleaseYear int      `json:"release-year"`
	Director    string   `json:"director"`
	Cast        []string `json:"cast"`
}

type Watcher struct {
	UserID              uuid.UUID `json:"user-id"`
	Username            string    `json:"username"`
	Email               string    `json:"email"`
	IsTrialSubscription bool      `json:"is-trial-subscription"`
	Movies              []Movie   `json:"movies"`
}

func ReadWatchers(filename string) ([]Watcher, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	} 

	defer file.Close()

	var watchers []Watcher

	err = json.NewDecoder(file).Decode(&watchers)
	if err != nil {
		return nil, err
	}

	return watchers, nil
}

func ToJson(watchers []Watcher) (string, error) {
	data, err := json.MarshalIndent(watchers, "", "	")
	if err != nil {
		return "", err
	}

	return string(data), nil
}