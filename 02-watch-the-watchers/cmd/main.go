package main

import (
	//"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/harshita-aggarwal/go-learning-collection/02-watch-the-watchers/watcher"
)

func main() {
	firstwatchers, err := watcher.ReadWatchers("./test_data/watchers.json")

	if err != nil {
		log.Fatal(err)
	}

	// popularMovies := watcher.FindPopularMovie(watchers)
	// data, err := json.MarshalIndent(popularMovies, "", "  ")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Popular Movies are: %v", string(data))
	// fmt.Println()

	// newWatcher := watcher.Watcher{
	// 	UserID: uuid.New(),
	// 	Username: "new.watcher",
	// 	Email: "new.watcher@example.com",
	// 	IsTrialSubscription: true,
	// }
	newWatcher2 := watcher.Watcher{
		UserID:              uuid.New(),
		Username:            "new.watcher2",
		Email:               "new.watcher2@example.com",
		IsTrialSubscription: true,
	}
	newWatcher3 := watcher.Watcher{
		UserID:              uuid.New(),
		Username:            "new.watcher3",
		Email:               "new.watcher3@example.com",
		IsTrialSubscription: true,
	}

	watcherIdOne, _ := uuid.Parse("1f2ad186-32c0-49d8-9bee-42f46215af2c")
	existingWatcher := watcher.Watcher{
		UserID:              watcherIdOne,
		Username:            "movie.lover.second.username",
		Email:               "movie.lover@example.com",
		IsTrialSubscription: true,
	}

	// if !(watcher.IsWatcherPresent(newWatcher.Email, watchers) || watcher.IsWatcherPresent(newWatcher.Username, watchers)) {
	// 	fmt.Println("Adding new watcher!!")
	// 	watchers = append(watchers, newWatcher)
	// }
	// if !(watcher.IsWatcherPresent(existingWatcher.Email, watchers) || watcher.IsWatcherPresent(existingWatcher.Username, watchers)) {
	// 	fmt.Println("Adding new watcher!")
	// 	watchers = append(watchers, existingWatcher)
	// }

	// for i, w := range watchers {
	// 	fmt.Printf("[%d of %d] Username: %s / Email: %s \n", i+1, len(watchers), w.Username, w.Email)
	// }

	// err = watcher.SaveJsonToFile(watchers, "./test_data/watchers-modified.json")

	// if err!=nil{
	// 	fmt.Println("Error Saving watchers: ", err)
	// }

	// fmt.Println("Watchers saved to file watchers-modified.json")

	secondwatchers := []watcher.Watcher{
		newWatcher2, newWatcher3, existingWatcher,
	}

	mergedWatchers := watcher.MergedWatcherSlices(firstwatchers, secondwatchers)

	for i, w := range mergedWatchers {
		fmt.Printf("[%d of %d] Username: %s / Email: %s \n", i+1, len(mergedWatchers), w.Username, w.Email)
	}
}
