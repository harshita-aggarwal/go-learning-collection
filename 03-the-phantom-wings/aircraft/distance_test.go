package aircraft

import (
	"os"
	"testing"
)

func TestBuildCache(t *testing.T) {
	type want struct {
		mapLength int
		err bool
		validation func (t *testing.T, distanceCache *DistanceCache)
	}
	testCases := []struct {
		name string
		jsonString string
		want want
	}{
		{
			name: "SINGLE_DISTANCE",
			jsonString: `[{
				"Origin": "A", "Destination": "B", "Distance": 100
			}]`,
			want: want{
				mapLength: 1,
				err: false,
				validation: func(t *testing.T, distanceCache *DistanceCache) {
					if len(distanceCache.Distances)!=1{
						t.Errorf("BuildCache() expected map length 1, got %v", len(distanceCache.Distances))
					}

					if distanceCache.Distances["A"]["B"] != 100 {
						t.Errorf("BuildCache() expected distance 100, got %v", distanceCache.Distances["A"]["B"])
					}
				},
			},
		},
		{
			name: "MULTIPLE_DISTANCES",
			jsonString: `[{
			"Origin":"A", "Destination":"B", "Distance":100}, {"Origin":"A", "Destination":"C", "Distance":302}]`,
			want: want{
				mapLength: 1,
				err: false,
				validation: func(t *testing.T, distanceCache *DistanceCache) {
					if len(distanceCache.Distances)!=1 {
						t.Errorf("BuildCache() expected map length 1, got %v", len(distanceCache.Distances))
					}

					if distanceCache.Distances["A"]["B"] != 100 {
						t.Errorf("BuildCache() expected distance 100, got %v", distanceCache.Distances["A"]["B"])
					}

					if distanceCache.Distances["A"]["C"] != 302 {
						t.Errorf("BuildCache() expected distance 302, got %v", distanceCache.Distances["A"]["C"])
					}
				},
			},
		},
		{
			name: "EMPTY_DISTANCE",
			jsonString: `[]`,
			want: want{
				mapLength: 0,
				err: false,
				validation: func(t *testing.T, distanceCache *DistanceCache){
					if len(distanceCache.Distances)!=0{
						t.Errorf("BuildCache() expected empty map, got %v", distanceCache.Distances)
					}
				},
			},
		},
		{
			name: "INVALID_JSON",
			jsonString: `[`,
			want: want{
				mapLength: 0,
				err: true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tmp, err := os.CreateTemp("", "distance_cache_*.json")

			if err != nil && tc.want.err{
				t.Errorf("Failed to create temp file: %v", err)
				return
			}

			defer os.Remove(tmp.Name())

			if _, err := tmp.Write([]byte(tc.jsonString)); err != nil{
				t.Errorf("Failed to write to file: %v", err)
				return
			}

			defer tmp.Close()

			distanceCache, err:= BuildCache(tmp.Name())

			if err!=nil && !tc.want.err {
				t.Errorf("BuildCache() error = %v, wanted error = %v", err, tc.want.err)
			}

			if err == nil && tc.want.err {
				t.Errorf("BuildCache() expected error, got nil")
			}

			if err == nil && (len(distanceCache.Distances) != tc.want.mapLength) {
				t.Errorf("BuildCache() expected map length %d, got %d", tc.want.mapLength, len(distanceCache.Distances))
			}

			if tc.want.validation != nil {
				tc.want.validation(t, distanceCache)
			}
		})
	}
}

func TestFindDistance(t *testing.T) {

	distanceCache := DistanceCache{
		Distances: map[string]map[string]int{
			"A":{
				"B": 100,
				"C": 302,
			},
			"B":{
				"C":300,
				"A": 120,
			},
			"C":{
				"A":220,
				"B":300,
			},
		},
	}

	testCases := []struct{
		name string
		origin string
		destination string
		want int
	}{
		{
			name: "	EXISTING_DISTANCE",
			origin: "A",
			destination: "B",
			want: 100,
		},
		{
			name: "	NON_EXISTING_ORIGIN",
			origin: "D",
			destination: "B",
			want: -1,
		},
		{
			name: "	NON_EXISTING_DESTINATION",
			origin: "B",
			destination: "D",
			want: -1,
		},
		{
			name: "	CASE_INSENSITIVE_LOOKUP",
			origin: "a",
			destination: "B",
			want: 100,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := FindDistance(&distanceCache, tc.origin, tc.destination)

			if got != tc.want {
				t.Errorf("expect %d, got: %d", tc.want, got)
			}
		})
	}
}