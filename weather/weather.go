package weather

import (
	"database/sql"
	"fmt"
	"math"

	"github.com/Giovanni299/Vulcano/database"
	planet "github.com/Giovanni299/Vulcano/planets"
)

const (
	sequia        = iota
	lluvia        = iota
	lluviaIntensa = iota
	optimo        = iota
	ninguno       = iota
)

var (
	sun = Coordinates{
		x: 0, y: 0,
	}
)

//DayResult struct to return informacion for a specific day.
type DayResult struct {
	Dia   int    `json:"dia"`
	Clima string `json:"clima"`
}

//GetWeather get the weather of the 10 years.
func GetWeather(db *sql.DB) ([]planet.WeatherResult, error) {
	Sequia := 0
	Lluvia := 0
	var LluviaIntensa string
	Optimo := 0
	lastWeather := -1

	valueDays := make([]string, 0, planet.FerengiDaysYear)
	valueDaysArgs := make([]interface{}, 0, planet.FerengiDaysYear*3)
	valueCoordinates := make([]string, 0, planet.FerengiDaysYear)
	valueCoorArgs := make([]interface{}, 0, planet.FerengiDaysYear*3)
	resultWeatherArgs := make([]interface{}, 0, 4)

	for day := uint(0); day < planet.FerengiDaysYear; day++ {
		coordinates := GetCoordinates(day)
		weather := calculateWeather(coordinates)
		//fmt.Printf("Day:%d Weather:%d Fer:(%.3f,%.3F) Beta:(%.3f,%.3F) Vul:(%.3f,%.3F)\n", i, weather, coordinates.coorFerengi.x, coordinates.coorFerengi.y, coordinates.coorBetasoide.x, coordinates.coorBetasoide.y, coordinates.coorVulcano.x, coordinates.coorVulcano.y)

		if weather != lastWeather {
			lastWeather = weather
			switch weather {
			case 0:
				Sequia++
			case 1:
				Lluvia++
			case 2:
				LluviaIntensa += fmt.Sprintf("%d,", day)
				lastWeather = lluvia
			case 3:
				Optimo++
			}
		}

		valueDays = append(valueDays, fmt.Sprintf("($%d, $%d)", day*2+1, day*2+2))
		valueDaysArgs = append(valueDaysArgs, day)
		valueDaysArgs = append(valueDaysArgs, weather)

		valueCoordinates = append(valueCoordinates, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d)", day*7+1, day*7+2, day*7+3, day*7+4, day*7+5, day*7+6, day*7+7))
		valueCoorArgs = append(valueCoorArgs, day)
		valueCoorArgs = append(valueCoorArgs, coordinates.coorBetasoide.x)
		valueCoorArgs = append(valueCoorArgs, coordinates.coorBetasoide.y)
		valueCoorArgs = append(valueCoorArgs, coordinates.coorFerengi.x)
		valueCoorArgs = append(valueCoorArgs, coordinates.coorFerengi.y)
		valueCoorArgs = append(valueCoorArgs, coordinates.coorVulcano.x)
		valueCoorArgs = append(valueCoorArgs, coordinates.coorVulcano.y)
	}

	resultWeatherArgs = append(resultWeatherArgs, Sequia)
	resultWeatherArgs = append(resultWeatherArgs, Lluvia)
	resultWeatherArgs = append(resultWeatherArgs, LluviaIntensa)
	resultWeatherArgs = append(resultWeatherArgs, Optimo)
	err := database.InsertData(db, valueDays, valueDaysArgs, valueCoordinates, valueCoorArgs, resultWeatherArgs)
	if err != nil {
		return nil, err
	}

	weatherResult := []planet.WeatherResult{
		{
			Weather: "Sequia", Value: fmt.Sprintf("(%d)", Sequia),
		},
		{
			Weather: "Lluvia", Value: fmt.Sprintf("(%d)", Lluvia),
		},
		{
			Weather: "LluviaIntensa", Value: LluviaIntensa,
		},
		{
			Weather: "Optimo", Value: fmt.Sprintf("(%d)", Optimo),
		},
	}

	println("Periodos de sequia: " + fmt.Sprint(Sequia))
	println("Periodos de Lluvia: " + fmt.Sprint(Lluvia))
	println("Periodos Optimo: " + fmt.Sprint(Optimo))
	fmt.Printf("Dias de lluvia intensa: %v\n", LluviaIntensa)

	return weatherResult, err
}

//GetDay return the result of the weather for a specific day.
func GetDay(db *sql.DB, day int) (DayResult, error) {
	dayResult, err := database.GetDay(db, day)
	if err != nil {
		return DayResult{}, err
	}

	weatherDay := "Clima no identificado"
	switch dayResult {
	case 0:
		weatherDay = "Sequia"
	case 1:
		weatherDay = "Lluvia"
	case 2:
		weatherDay = "Lluvia Intensa"
	case 3:
		weatherDay = "Condiciones Optimas"
	}

	return DayResult{
		Dia:   day,
		Clima: weatherDay,
	}, nil
}

//calculateWeather using the locations of the planets.
func calculateWeather(coordinates CoordinatesPlanet) int {
	isAligned, weather := planetsAligned(coordinates)
	if isAligned {
		return weather
	}

	return sunInTriangle(coordinates)
}

//planetsAligned validate if the 3 or 4 satellites are aligned.
func planetsAligned(coor CoordinatesPlanet) (bool, int) {
	//If x or y in the 3 point are equals, the planets are aligned.
	if (coor.coorBetasoide.x == coor.coorFerengi.x && coor.coorBetasoide.x == coor.coorVulcano.x) ||
		(coor.coorBetasoide.y == coor.coorFerengi.y && coor.coorBetasoide.y == coor.coorVulcano.y) {
		//Validate if it match es the sun.
		if (coor.coorVulcano.x == sun.x) || (coor.coorVulcano.y == sun.y) {
			return true, sequia
		}

		return true, optimo
	}

	//First, validate the 3 points of the planets.
	val1 := (coor.coorBetasoide.y - coor.coorFerengi.y) / (coor.coorBetasoide.x - coor.coorFerengi.x)
	val2 := (coor.coorVulcano.y - coor.coorBetasoide.y) / (coor.coorVulcano.x - coor.coorBetasoide.x)

	//if val1 == val2 {
	if val1/val2 >= 0.97 && val1/val2 < 1.03 {
		//After, validate the 2 points of the planets and the ponit of the sun.
		val3 := (coor.coorBetasoide.y - sun.y) / (coor.coorBetasoide.x - sun.x)
		if val3 == val1 {
			return true, sequia
		}

		return true, optimo
	}

	return false, 0
}

//sunInTriangle validate if the planets shape a triangle and the sun is inside.
func sunInTriangle(coor CoordinatesPlanet) int {
	var s = coor.coorFerengi.y*coor.coorVulcano.x - coor.coorFerengi.x*coor.coorVulcano.y + (coor.coorVulcano.y-coor.coorFerengi.y)*sun.x + (coor.coorFerengi.x-coor.coorVulcano.x)*sun.y
	var t = coor.coorFerengi.x*coor.coorBetasoide.y - coor.coorFerengi.y*coor.coorBetasoide.x + (coor.coorFerengi.y-coor.coorBetasoide.y)*sun.x + (coor.coorBetasoide.x-coor.coorFerengi.x)*sun.y
	if (s <= 0) != (t <= 0) {
		return ninguno
	}

	var A = -coor.coorBetasoide.y*coor.coorVulcano.x + coor.coorFerengi.y*(coor.coorVulcano.x-coor.coorBetasoide.x) + coor.coorFerengi.x*(coor.coorBetasoide.y-coor.coorVulcano.y) + coor.coorBetasoide.x*coor.coorVulcano.y
	if A < 0 {
		s = -s
		t = -t
		A = -A
	}

	if s >= 0 && t >= 0 && (s+t) <= A {
		perimeter := calculatePerimeter(coor)
		if perimeter >= planet.MaxPerimeterTriangle {
			//println("Perimeter: " + fmt.Sprintf("%f", perimeter))
			return lluviaIntensa
		}

		return lluvia
	}

	return ninguno
}

//calculatePerimeter of actual triangle.
func calculatePerimeter(coor CoordinatesPlanet) float64 {
	distFeBe := math.Sqrt(math.Pow((coor.coorBetasoide.x-coor.coorFerengi.x), 2) + math.Pow((coor.coorBetasoide.y-coor.coorFerengi.y), 2))
	distFeVu := math.Sqrt(math.Pow((coor.coorVulcano.x-coor.coorFerengi.x), 2) + math.Pow((coor.coorVulcano.y-coor.coorFerengi.y), 2))
	distVuBe := math.Sqrt(math.Pow((coor.coorBetasoide.x-coor.coorVulcano.x), 2) + math.Pow((coor.coorBetasoide.y-coor.coorVulcano.y), 2))

	return distFeBe + distFeVu + distVuBe
}
