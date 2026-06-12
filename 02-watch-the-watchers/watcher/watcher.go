package watcher

import (
	"encoding/json"
	"os"
	"sort"
	"strings"

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

var allMovies = make(map[string]Movie)
var moviePopularity = make(map[string]int)

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

func normalizeTitles(title string) string {
	words := strings.Fields(title)
	normalized := strings.ToLower(strings.Join(words, "-"))
	return normalized
}

func populateBothMaps(watchers []Watcher) {
	for _, watcher := range watchers {
		for _, movie := range watcher.Movies {
			movieTitle := normalizeTitles(movie.Title)
			allMovies[movieTitle] = movie
			moviePopularity[movieTitle]++
		}
	}
} 

func sortMovies(movies []Movie) []Movie {
	sort.Slice(movies, func(i, j int) bool {
		titleI := normalizeTitles(movies[i].Title)
		titleJ := normalizeTitles(movies[j].Title)

		if titleI == titleJ {
			return movies[i].ReleaseYear < movies[j].ReleaseYear
		}

		return titleI < titleJ
	})

	return movies
}

func FindPopularMovie(watchers []Watcher) []Movie {
	populateBothMaps(watchers)

	var popularMovies = []Movie{}

	for normalizeTitle, count := range moviePopularity {
		if count >1 {
			popularMovies = append(popularMovies, allMovies[normalizeTitle])
		}
	}

	return sortMovies(popularMovies)

}