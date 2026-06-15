package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/harshita-aggarwal/go-learning-collection/03-the-phantom-wings/aircraft"
)



func main(){

	/*----------------------*/
	/*AIRCRAFT STRUCT OBJECT*/
	/*----------------------*/

	aircraftIdOne, _ := uuid.Parse("4bc8e887-c2ba-46d6-8ecb-ce175775e3f9")
	newAircraft := aircraft.Aircraft{
		AircraftId: aircraftIdOne,
		RegistrationCode: "newAircraft1",
		CurrentMileage: 0,
	} 

	/*------------------*/
	/*FOR STRUCT TO JSON*/
	/*------------------*/
	fmt.Println("Aircraft One -- Initial")
	aircraft.StructToJson(newAircraft)
	
	/*------------------*/
	/*FOR ADDING MILEAGE*/
	/*------------------*/
	aircraft.AddMileage(&newAircraft, 400)
	aircraft.StructToJson(newAircraft)
}
