package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/harshita-aggarwal/go-learning-collection/03-the-phantom-wings/aircraft"
)



func main(){

	/*----------------------*/
	/*AIRLINE STRUCT OBJECT*/
	/*----------------------*/
	airlineOne := aircraft.Airline{
		AirlineId: uuid.New(),
		Name: "Airline One",
		OriginCountry: "India",
		IATACode: "AI",
	}

	/*----------------------*/
	/*AIRCRAFT STRUCT OBJECT*/
	/*----------------------*/
	aircraftOne := aircraft.Aircraft{
		AircraftId: uuid.New(),
		RegistrationCode: "aircraft1",
		CurrentMileage: 0,
		Airline: &airlineOne,
	} 
	aircraftTwo := aircraft.Aircraft{
		AircraftId: uuid.New(),
		RegistrationCode: "aircraft2",
		CurrentMileage: 100,
		Airline: &airlineOne,
	} 

	/*------------------*/
	/*FOR STRUCT TO JSON*/
	/*------------------*/
	fmt.Println("Aircraft One -- Initial")
	aircraft.StructToJson(aircraftOne)
	fmt.Println("Aircraft Two -- Initial")
	aircraft.StructToJson(aircraftTwo)
	
	
	/*------------------*/
	/*FOR ADDING MILEAGE*/
	/*------------------*/
	aircraft.AddMileage(&aircraftOne, 400)
	aircraft.AddMileage(&aircraftTwo, 400)

	/*---------------------*/
	/*CHANGING AIRLINE DATA*/
	/*---------------------*/
	airlineOne.IATACode = "AIR"

	fmt.Println("Aircraft One -- After Airline Update")
	aircraft.StructToJson(aircraftOne)
	fmt.Println("Aircraft Two -- After Airline Update")
	aircraft.StructToJson(aircraftTwo)
}
