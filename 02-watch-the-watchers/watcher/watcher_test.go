package watcher

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

var (
	movieOne = Movie{
		Title:       "The Recipe For Chaos",
		Genres:      []string{"Comedy", "Adventure"},
		ReleaseYear: 2021,
		Director:    "Caldwell Dany",
		Cast:        []string{"Olivia Brown", "Emily Carter", "Michael Johnson", "Lucas Hall"},
	}

	movieTwo = Movie{
		Title:       "48 Hours With Alice",
		Genres:      []string{"Fantasy", "Sci-Fi", "Action"},
		ReleaseYear: 2022,
		Director:    "Benjamin Moore",
		Cast:        []string{"Sophia Wilson", "Ethan Cartez", "Isabella Anderson"},
	}

	movieThree = Movie{
		Title:       "The First Unknown",
		Genres:      []string{"Thriller", "Action"},
		ReleaseYear: 2023,
		Director:    "Amelia Young",
		Cast:        []string{"Alexander Walker", "Lucas Hall", "Mia Harris", "Henry King", "Isabella Anderson"},
	}

	movieFour = Movie{
		Title:       "White Wars",
		Genres:      []string{"Action", "Sci-Fi"},
		ReleaseYear: 2024,
		Director:    "Benjamin Moore",
		Cast:        []string{"Harper Wright", "Charlotte Lee", "Samuel Carter", "Grace White", "Evelyn Phillips"},
	}

	movieFive = Movie{
		Title:       "Binary",
		Genres:      []string{"Thriller", "Mystery"},
		ReleaseYear: 2021,
		Director:    "Caroline Young",
		Cast:        []string{"Darla Alexander", "Billy Evans", "Pedro Hall"},
	}

	watcherIdOne, _   = uuid.Parse("1f2ad186-32c0-49d8-9bee-42f46215af2c")
	watcherIdTwo, _   = uuid.Parse("2a817864-3683-4431-b1cd-dcd395aef147")
	watcherIdThree, _ = uuid.Parse("3bee74a8-b06c-4a3c-be2e-4b2f21c80f40")
	watcherIdFour, _  = uuid.Parse("4a1e71e7-018b-4f9e-b241-96e39d2051f6")

	watcherOne = Watcher{
		UserID:              watcherIdOne,
		Username:            "movie.lover",
		Email:               "movie.lover@example.com",
		IsTrialSubscription: true,
		Movies:              []Movie{movieOne, movieTwo, movieFive},
	}

	watcherTwo = Watcher{
		UserID:              watcherIdTwo,
		Username:            "film.fanatic",
		Email:               "film.fanatic@example.com",
		IsTrialSubscription: false,
		Movies:              []Movie{movieFour},
	}

	watcherThree = Watcher{
		UserID:              watcherIdThree,
		Username:            "cinema.critic",
		Email:               "cinema.critic@example.com",
		IsTrialSubscription: false,
		Movies:              []Movie{movieFour, movieOne, movieThree},
	}

	watcherFour = Watcher{
		UserID:              watcherIdFour,
		Username:            "movie.always",
		Email:               "movie.always@example.com",
		IsTrialSubscription: false,
		Movies:              []Movie{movieTwo, movieFour},
	}
)

func TestLoadWatchersFromFile(t *testing.T) {
	type testCase struct {
		name     string
		filePath string
		want     []Watcher
		wantErr  bool
	}

	wantWatchers := []Watcher{watcherOne, watcherTwo, watcherThree, watcherFour}

	testCases := []testCase{
		{
			name:     "VALID_JSON_FILE",
			filePath: "../test_data/watchers.json",
			want:     wantWatchers,
			wantErr:  false,
		},
		{
			name:     "JSON_FILE_NOT_EXIST",
			filePath: "../test_data/non_existent_file.json",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "INVALID_JSON_FILE",
			filePath: "../test_data/watchers_contains_error.json",
			want:     nil,
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ReadWatchers(tc.filePath)

			if err == nil && tc.wantErr {
				t.Errorf("%v Expected error, got nil", tc.name)
				return
			}

			if err != nil && !tc.wantErr {
				t.Errorf("%v Unexpected error: %v", tc.name, err)
				return
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("%v Expected %v, got %v", tc.name, tc.want, got)
				return
			}
		})
	}
}

func TestNormalizeTitle(t *testing.T) {
	testCases := []struct {
		name     string
		original string
		want     string
	}{
		{
			name:     "LEADING_AND_TRAILING_SPACES",
			original: "   The Matrix   ",
			want:     "the-matrix",
		},
		{
			name:     "MULTIPLE_SPACES",
			original: "The   Matrix     Two",
			want:     "the-matrix-two",
		},
		{
			name:     "MIXED_CASE",
			original: "tHe MaTrIx tHREe",
			want:     "the-matrix-three",
		},
		{
			name:     "NO_SPACES",
			original: "TheMatrixFour",
			want:     "thematrixfour",
		},
		{
			name:     "PRESERVE_SPECIAL_CHARACTERS",
			original: "The Matrix Five !@#",
			want:     "the-matrix-five-!@#",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := normalizeTitles(tc.original)

			if got != tc.want {
				t.Errorf("%v Expected %v, got %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestPopulateAllMovies(t *testing.T) {
	testCases := []struct {
		name     string
		watchers []Watcher
		want     map[string]Movie
	}{
		{
			name:     "NO_WATCHERS",
			watchers: []Watcher{},
			want:     map[string]Movie{},
		},
		{
			name: "SINGLE_WATCHER_SINGLE_MOVIE",
			watchers: []Watcher{
				{
					UserID:              watcherIdOne,
					Username:            "single.movie",
					Email:               "single.movie@example.com",
					IsTrialSubscription: true,
					Movies:              []Movie{movieOne},
				},
			},
			want: map[string]Movie{
				"the-recipe-for-chaos": movieOne,
			},
		},
		{
			name: "MULTIPLE_WATCHERS_MULTIPLE_MOVIES",
			watchers: []Watcher{
				watcherOne,
				watcherTwo,
				watcherThree,
				watcherFour,
			},
			want: map[string]Movie{
				"the-recipe-for-chaos": movieOne,
				"48-hours-with-alice":  movieTwo,
				"the-first-unknown":    movieThree,
				"white-wars":           movieFour,
				"binary":               movieFive,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			allMovies = make(map[string]Movie)

			populateBothMaps(tc.watchers)

			if !reflect.DeepEqual(allMovies, tc.want) {
				t.Errorf("%v Expected %v, got %v", tc.name, tc.want, allMovies)
			}
		})
	}
}

func TestFindPopularMovies(t *testing.T) {
	testCases := []struct {
		name     string
		watchers []Watcher
		want     []Movie
	}{
		{
			name:     "NO_WATCHERS",
			watchers: []Watcher{},
			want:     []Movie{},
		},
		{
			name: "NO_POPULAR_MOVIES",
			watchers: []Watcher{
				{
					UserID:              watcherIdOne,
					Username:            "single.movie",
					Email:               "single.movie@example.com",
					IsTrialSubscription: true,
					Movies:              []Movie{movieOne},
				},
				{
					UserID:              watcherIdTwo,
					Username:            "single.movie.2",
					Email:               "single.movie.2@example.com",
					IsTrialSubscription: false,
					Movies:              []Movie{movieTwo, movieThree},
				},
				{
					UserID:              watcherIdThree,
					Username:            "single.movie.3",
					Email:               "single.movie.3@example.com",
					IsTrialSubscription: false,
					Movies:              []Movie{movieFour, movieFive},
				},
			},
			want: []Movie{},
		},
		{
			name: "SINGLE_POPULAR_MOVIE",
			watchers: []Watcher{
				watcherOne,
				watcherThree,
			},
			want: []Movie{movieOne},
		},
		{
			name: "MULTIPLE_POPULAR_MOVIES",
			watchers: []Watcher{
				watcherOne,
				watcherTwo,
				watcherThree,
				watcherFour,
			},
			want: []Movie{movieTwo, movieOne, movieFour},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			allMovies = make(map[string]Movie)
			moviePopularity = make(map[string]int)

			//populateBothMaps(tc.watchers)

			got := FindPopularMovie(tc.watchers)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("%v Expected %v, got %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestSortMovies(t *testing.T) {
	movieTitleA := Movie{
		Title:       "The Movie A",
		Genres:      []string{"Comedy"},
		ReleaseYear: 2025,
		Director:    "Director A",
		Cast:        []string{"Actor A"},
	}

	movieTitleB := Movie{
		Title:       "The Movie B",
		Genres:      []string{"Drama"},
		ReleaseYear: 2025,
		Director:    "Director B",
		Cast:        []string{"Actor B"},
	}

	movieTitleC := Movie{
		Title:       "The Movie C",
		Genres:      []string{"Action"},
		ReleaseYear: 2025,
		Director:    "Director C",
		Cast:        []string{"Actor C"},
	}

	movieTitleXReleaseYear2024 := Movie{
		Title:       "The Same Title",
		Genres:      []string{"Drama"},
		ReleaseYear: 2024,
		Director:    "Director A",
		Cast:        []string{"Actor A"},
	}

	movieTitleXReleaseYear2021 := Movie{
		Title:       "tHe SaME tiTle   ",
		Genres:      []string{"Drama"},
		ReleaseYear: 2021,
		Director:    "Director B",
		Cast:        []string{"Actor B"},
	}

	testCases := []struct {
		name   string
		movies []Movie
		want   []Movie
	}{
		{
			name:   "NO_MOVIES",
			movies: []Movie{},
			want:   []Movie{},
		},
		{
			name: "SINGLE_MOVIE",
			movies: []Movie{
				movieOne,
			},
			want: []Movie{
				movieOne,
			},
		},
		{
			name: "MULTIPLE_MOVIES_DIFFERENT_TITLES",
			movies: []Movie{
				movieTitleB,
				movieTitleC,
				movieTitleA,
			},
			want: []Movie{
				movieTitleA,
				movieTitleB,
				movieTitleC,
			},
		},
		{
			name: "MULTIPLE_MOVIES_SAME_TITLE_DIFFERENT_YEARS",
			movies: []Movie{
				movieTitleXReleaseYear2024,
				movieTitleXReleaseYear2021,
			},
			want: []Movie{
				movieTitleXReleaseYear2021,
				movieTitleXReleaseYear2024,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := sortMovies(tc.movies)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("%v Expected %v, got %v", tc.name, tc.want, got)
			}
		})
	}
}