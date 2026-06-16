package main

import (
	"fmt"
	"log"

	//"github.com/google/uuid"
	"github.com/harshita-aggarwal/go-learning-collection/03-the-phantom-wings/aircraft"
)



func main(){


	// /*----------------------*/
	// /*AIRLINE STRUCT OBJECT*/
	// /*----------------------*/
	// airlineOne := aircraft.Airline{
	// 	AirlineId: uuid.New(),
	// 	Name: "Airline One",
	// 	OriginCountry: "India",
	// 	IATACode: "AI",
	// }

	// /*----------------------*/
	// /*AIRCRAFT STRUCT OBJECT*/
	// /*----------------------*/
	// aircraftOne := aircraft.Aircraft{
	// 	AircraftId: uuid.New(),
	// 	RegistrationCode: "aircraft1",
	// 	CurrentMileage: 0,
	// 	Airline: &airlineOne,
	// } 
	// aircraftTwo := aircraft.Aircraft{
	// 	AircraftId: uuid.New(),
	// 	RegistrationCode: "aircraft2",
	// 	CurrentMileage: 100,
	// 	Airline: &airlineOne,
	// } 

	// /*------------------*/
	// /*FOR STRUCT TO JSON*/
	// /*------------------*/
	// fmt.Println("Aircraft One -- Initial")
	// aircraft.StructToJson(aircraftOne)
	// fmt.Println("Aircraft Two -- Initial")
	// aircraft.StructToJson(aircraftTwo)
	
	
	// /*------------------*/
	// /*FOR ADDING MILEAGE*/
	// /*------------------*/
	// aircraft.AddMileage(&aircraftOne, 400)
	// aircraft.AddMileage(&aircraftTwo, 400)

	// /*---------------------*/
	// /*CHANGING AIRLINE DATA*/
	// /*---------------------*/
	// airlineOne.IATACode = "AIR"

	// fmt.Println("Aircraft One -- After Airline Update")
	// aircraft.StructToJson(aircraftOne)
	// fmt.Println("Aircraft Two -- After Airline Update")
	// aircraft.StructToJson(aircraftTwo)

	/*-------------*/
	/*BUIDING CACHE*/
	/*-------------*/
	distanceCache, err := aircraft.BuildCache("./test_data/distances.json")

	if err != nil{
		log.Fatal("Error building cache: ", err)
	}
	
	/*---------------------------*/
	/*TO PRINT DISTANCE IF EXISTS*/
	/*---------------------------*/
	distance := aircraft.FindDistance(distanceCache, "London", "Paris")
	if distance != -1 {
		fmt.Println(distance)
	}
	fmt.Println()
	distance = aircraft.FindDistance(distanceCache, "JAKARta", "singapore")
	if distance != -1 {
		fmt.Println(distance)
	}
}
