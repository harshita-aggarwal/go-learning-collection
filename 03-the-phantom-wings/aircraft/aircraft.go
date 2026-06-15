package aircraft

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type Aircraft struct{
	AircraftId uuid.UUID `json:"aircraft-id,omitempty"`
	RegistrationCode string `json:"registration-code,omitempty"`
	CurrentMileage int `json:"current-mileage,omitempty"`
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