package weather

import (
	"math"

	planet "github.com/Giovanni299/Vulcano/planets"
)

// DegreesCircumference specifies the number of degrees in a whole turn of a circle.
const DegreesCircumference = 360.0

//Coordinates coordiante system in the Cartesian plane.
type Coordinates struct {
	x float64
	y float64
}

//CoordinatesPlanet of Planets by day.
type CoordinatesPlanet struct {
	coorFerengi   Coordinates
	coorBetasoide Coordinates
	coorVulcano   Coordinates
}

//GetCoordinates return coordinates information by day.
func GetCoordinates(day uint) CoordinatesPlanet {
	coordinatesPlanet := CoordinatesPlanet{
		coorFerengi:   calculateCoordinates(day, planet.Planets[0]),
		coorBetasoide: calculateCoordinates(day, planet.Planets[1]),
		coorVulcano:   calculateCoordinates(day, planet.Planets[2]),
	}

	return coordinatesPlanet
}

//calculateCoordinates get coordinates by each planet.
func calculateCoordinates(day uint, planet planet.Planet) Coordinates {
	var degrees float64
	if planet.HoraryDirection {
		degrees = float64(DegreesCircumference - (day * planet.AngularSpeed % DegreesCircumference))

	} else {
		degrees = float64(day * planet.AngularSpeed % DegreesCircumference)
	}

	coordinate := Coordinates{
		x: math.Round(planet.OrbitRadius*math.Cos(degrees*math.Pi/180)*1000) / 1000,
		y: math.Round(planet.OrbitRadius*math.Sin(degrees*math.Pi/180)*1000) / 1000,
	}

	return coordinate
}
