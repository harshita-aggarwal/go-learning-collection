package aircraft

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type Airline struct{
	AirlineId uuid.UUID `json:"airline-id,omitempty"`
	Name string `json:"name,omitempty"`
	OriginCountry string `json:"origin-country,omitempty"`
	IATACode string `json:"IATA-code,omitempty"`
}

type Aircraft struct{
	AircraftId uuid.UUID `json:"aircraft-id,omitempty"`
	RegistrationCode string `json:"registration-code,omitempty"`
	CurrentMileage int `json:"current-mileage,omitempty"`
	Airline *Airline `json:"airline,omitempty"`
}

func StructToJson(aircraft Aircraft) {
	data, err := json.MarshalIndent(aircraft, "", "	")

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshalling aircraft to JSON, ", err)
	}

	fmt.Println(string(data))
}

func AddMileage(pointerToAircraft *Aircraft, mileage int) {
	pointerToAircraft.CurrentMileage += mileage
}
