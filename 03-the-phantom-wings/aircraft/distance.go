package aircraft

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Distance struct {
	Origin string `json:"origin,omitempty"`
	Destination string `json:"destination,omitempty"`
	Distance int `json:"distance,omitempty"`
}

type DistanceCache struct{
	Distances map[string]map[string]int
}

func BuildCache(filename string) (*DistanceCache, error) {
	
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var distancesFromFile = []Distance{}

	err = json.NewDecoder(file).Decode(&distancesFromFile)
	if err != nil {
		return nil, err
	}

	distances := make(map[string]map[string]int)

	for _, distance := range distancesFromFile{
		origin := Normalize(distance.Origin)
		destination := Normalize(distance.Destination)

		if _, ok := distances[origin]; !ok {
			distances[origin] = make(map[string]int)
		}
		distances[origin][destination] = distance.Distance
	}
	return &DistanceCache{Distances: distances,}, nil
}

func Normalize(location string) string {
	return strings.ToUpper(strings.TrimSpace(location))
}


func FindDistance(distanceCache *DistanceCache, origin, destination string) (int){
	o := Normalize(origin)
	d := Normalize(destination)

	if _, ok := distanceCache.Distances[o]; !ok{
		fmt.Printf("Route Between %s And %s Doesn't Exist", o, d)
		return -1
	}

	if distance, ok := distanceCache.Distances[o][d]; ok{
		fmt.Printf("Distance between %s And %s: ", o, d)
		return distance
	}

	fmt.Println("No Such Origin")
	return -1
}
