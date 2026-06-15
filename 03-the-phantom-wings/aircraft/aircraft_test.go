package aircraft

import (
	"reflect"
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

func TestAirlineChangesReflection(t *testing.T) {
	dummyUUID := uuid.New()

	testCases := []struct{
		name string
		airline Airline
		want Airline
	}{
		{
			name: "CHANGE_AIRLINE_NAME",
			airline: Airline{
				Name: "Airline A",
				AirlineId: dummyUUID,
				OriginCountry: "Country A",
				IATACode: "AA",
			},
			want: Airline{
				Name: "Airline B",
				AirlineId: dummyUUID,
				OriginCountry: "Country A",
				IATACode: "AA",
			},
		},
		{
			name: "CHANGE_AIRLINE_ID",
			airline: Airline{
				Name: "Airline A",
				AirlineId: dummyUUID,
				OriginCountry: "Country A",
				IATACode: "AA",
			},
			want: Airline{
				Name: "Airline A",
				AirlineId: uuid.New(),
				OriginCountry: "Country A",
				IATACode: "AA",
			},
		},
		{
			name: "CHANGE_AIRLINE_CODE",
			airline: Airline{
				Name: "Airline A",
				AirlineId: dummyUUID,
				OriginCountry: "Country A",
				IATACode: "AA",
			},
			want: Airline{
				Name: "Airline A",
				AirlineId: dummyUUID,
				OriginCountry: "Country A",
				IATACode: "AB",
			},
		},
		{
			name: "CHANGE_AIRLINE_COUNTRY",
			airline: Airline{
				Name: "Airline A",
				AirlineId: dummyUUID,
				OriginCountry: "Country A",
				IATACode: "AA",
			},
			want: Airline{
				Name: "Airline A",
				AirlineId: dummyUUID,
				OriginCountry: "Country B",
				IATACode: "AA",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			aircraft := Aircraft{
				Airline: &tc.airline,
			}
			tc.airline.IATACode = tc.want.IATACode
			tc.airline.AirlineId = tc.want.AirlineId
			tc.airline.Name = tc.want.Name
			tc.airline.OriginCountry = tc.want.OriginCountry

			if !reflect.DeepEqual(aircraft.Airline, &tc.airline) {
				t.Errorf("expected: %v, got: %v", tc.airline, aircraft.Airline)
			}
		})
	}
}