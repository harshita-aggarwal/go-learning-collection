package main

import (
	"fmt"
	"log"

	"github.com/harshita-aggarwal/go-learning-collection/02-watch-the-watchers/watcher"
)

func main() {
	watchers, err := watcher.ReadWatchers("../test_data/watchers.json")

	if err != nil {
		log.Fatal(err)
	}

	jsonString, err := watcher.ToJson(watchers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(jsonString)
}
