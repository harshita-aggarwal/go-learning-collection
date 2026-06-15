package aircraft

import (
	"testing"
	"github.com/google/uuid"
)

func TestAddMileage(t *testing.T){
	testCases := []struct{
		name string
		mileage int
		toAdd int
		want int
	}{
		{
			name: "ADDING_POSITIVE_MILEAGE",
			mileage: 500,
			toAdd: 400,
			want: 900,
		},
		{
			name: "ADDING_ZERO_MILEAGE",
			mileage: 500,
			toAdd: 0,
			want: 500,
		},
		{
			name: "INITIAL_ZERO_MILEAGE",
			mileage: 0,
			toAdd: 500,
			want: 500,
		},
	}

	for _, tc := range testCases{
		t.Run(tc.name, func(t *testing.T) {
			aircraft := Aircraft{
				AircraftId: uuid.New(),
				RegistrationCode: "some-aircraft-no.",
				CurrentMileage: tc.mileage,
			}
			AddMileage(&aircraft, tc.toAdd)
			got := aircraft.CurrentMileage

			if got!= tc.want {
				t.Errorf("Expected %v, got %v", tc.want, got)
			}
		})
	}
}