package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/harshita-aggarwal/go-learning-collection/02-watch-the-watchers/watcher"
)

func main() {
	watchers, err := watcher.ReadWatchers("./test_data/watchers.json")

	if err != nil {
		log.Fatal(err)
	}

	popularMovies := watcher.FindPopularMovie(watchers)
	data, err := json.MarshalIndent(popularMovies, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Popular Movies are: %v", string(data))
}
