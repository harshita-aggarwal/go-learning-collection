package main

import (
	//"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/harshita-aggarwal/go-learning-collection/02-watch-the-watchers/watcher"
)

func main() {

	/*----------------------------*/
	/*TO READ THE JSON FILE*/
	/*----------------------------*/
	firstwatchers, err := watcher.ReadWatchers("./test_data/watchers.json")

	if err != nil {
		log.Fatal(err)
	}

	/*----------------------------*/
	/*TO FIND POPULAR MOVIES*/
	/*----------------------------*/
	// popularMovies := watcher.FindPopularMovie(firstwatchers)
	// data, err := json.MarshalIndent(popularMovies, "", "  ")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Popular Movies are: %v", string(data))
	// fmt.Println()

	/*----------------------------*/
	/*NEW AND EXISTING WATCHERS FOR INPUT*/
	/*----------------------------*/
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
	
	/*----------------------------*/
	/*TO ADD NEW WATCHER IF NOT PRESENT*/
	/*----------------------------*/
	// if !(watcher.IsWatcherPresent(newWatcher.Email, firstwatchers) || watcher.IsWatcherPresent(newWatcher.Username, firstwatchers)) {
	// 	fmt.Println("Adding new watcher!!")
	// 	firstwatchers = append(firstwatchers, newWatcher)
	// }
	// if !(watcher.IsWatcherPresent(existingWatcher.Email, firstwatchers) || watcher.IsWatcherPresent(existingWatcher.Username, firstwatchers)) {
	// 	fmt.Println("Adding new watcher!")
	// 	firstwatchers = append(firstwatchers, existingWatcher)
	// }

	/*----------------------------*/
	/*TO SAVE JSON TO FILE*/
	/*----------------------------*/
	// err = watcher.SaveJsonToFile(watchers, "./test_data/watchers-modified.json")
	
	// if err!=nil{
		// 	fmt.Println("Error Saving watchers: ", err)
		// }
		
	// fmt.Println("Watchers saved to file watchers-modified.json")
		
	/*----------------------------*/
	/*TO MERGE TWO WATCHER SLICES*/
	/*----------------------------*/
	secondwatchers := []watcher.Watcher{
		newWatcher2, newWatcher3, existingWatcher,
	}

	mergedWatchers := watcher.MergedWatcherSlices(firstwatchers, secondwatchers)

	for i, w := range mergedWatchers {
		fmt.Printf("[%d of %d] Username: %s / Email: %s \n", i+1, len(mergedWatchers), w.Username, w.Email)
	}
}
