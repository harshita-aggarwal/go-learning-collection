package watcher

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

var (
	movieRecipeForChaos = Movie{
		Title:       "The Recipe For Chaos",
		Genres:      []string{"Comedy", "Adventure"},
		ReleaseYear: 2021,
		Director:    "Caldwell Dany",
		Cast:        []string{"Olivia Brown", "Emily Carter", "Michael Johnson", "Lucas Hall"},
	}

	movieAlice = Movie{
		Title:       "48 Hours With Alice",
		Genres:      []string{"Fantasy", "Sci-Fi", "Action"},
		ReleaseYear: 2022,
		Director:    "Benjamin Moore",
		Cast:        []string{"Sophia Wilson", "Ethan Cartez", "Isabella Anderson"},
	}

	movieBinary = Movie{
		Title:       "Binary",
		Genres:      []string{"Thriller", "Mystery"},
		ReleaseYear: 2021,
		Director:    "Caroline Young",
		Cast:        []string{"Darla Alexander", "Billy Evans", "Pedro Hall"},
	}

	movieWhiteWars = Movie{
		Title:       "White Wars",
		Genres:      []string{"Action", "Sci-Fi"},
		ReleaseYear: 2024,
		Director:    "Benjamin Moore",
		Cast:        []string{"Harper Wright", "Charlotte Lee", "Samuel Carter", "Grace White", "Evelyn Phillips"},
	}

	movieFirstUnknown = Movie{
		Title:       "The First Unknown",
		Genres:      []string{"Thriller", "Action"},
		ReleaseYear: 2023,
		Director:    "Amelia Young",
		Cast:        []string{"Alexander Walker", "Lucas Hall", "Mia Harris", "Henry King", "Isabella Anderson"},
	}
)

var (
	watcherOne = Watcher{
		UserID:              uuid.MustParse("1f2ad186-32c0-49d8-9bee-42f46215af2c"),
		Username:            "movie.lover",
		Email:               "movie.lover@example.com",
		IsTrialSubscription: true,
		Movies: []Movie{
			movieRecipeForChaos,
			movieAlice,
			movieBinary,
		},
	}

	watcherTwo = Watcher{
		UserID:              uuid.MustParse("2a817864-3683-4431-b1cd-dcd395aef147"),
		Username:            "film.fanatic",
		Email:               "film.fanatic@example.com",
		IsTrialSubscription: false,
		Movies: []Movie{
			movieWhiteWars,
		},
	}

	watcherThree = Watcher{
		UserID:              uuid.MustParse("3bee74a8-b06c-4a3c-be2e-4b2f21c80f40"),
		Username:            "cinema.critic",
		Email:               "cinema.critic@example.com",
		IsTrialSubscription: false,
		Movies: []Movie{
			movieWhiteWars,
			movieRecipeForChaos,
			movieFirstUnknown,
		},
	}

	watcherFour = Watcher{
		UserID:              uuid.MustParse("4a1e71e7-018b-4f9e-b241-96e39d2051f6"),
		Username:            "movie.always",
		Email:               "movie.always@example.com",
		IsTrialSubscription: false,
		Movies: []Movie{
			movieAlice,
			movieWhiteWars,
		},
	}
)
func TestReadWatcher(t *testing.T) {
	t.Run("read watchers from valid JSON file", func(t *testing.T) {
		got, err := ReadWatchers("../test_data/watchers.json")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	
		want := []Watcher{
		watcherOne, 
		watcherTwo,
		watcherThree,
		watcherFour,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v, want: %+v", got, want)
	}


	})


	t.Run("returns error when file doesn't exist", func(t *testing.T) {
		_, err := ReadWatchers("../test_data/watchers-does-not-exist.json")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	} )
	t.Run("returns error when file contains error", func(t *testing.T) {
		_, err := ReadWatchers("../test_data/watchers-cotains-error.json")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	} )


}