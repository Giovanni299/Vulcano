package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	planet "github.com/Giovanni299/Vulcano/planets"
	"github.com/labstack/echo"

	"github.com/Giovanni299/Vulcano/database"
	"github.com/Giovanni299/Vulcano/weather"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//WeatherService path weather.
const WeatherService = "/weather"

const DayService = "/clima"

var (
	db            *sql.DB
	weatherResult []planet.WeatherResult
	dbHost        string
	dbUsername    string
	dbName        string
	dbPort        string
	dbPassword    string
)

func init() {
	//Load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file.")
	}

	dbHost = os.Getenv("dbHost")
	dbUsername = os.Getenv("dbUsername")
	dbPassword = os.Getenv("dbPassword")
	dbName = os.Getenv("dbName")
	dbPort = os.Getenv("dbPort")
}

func main() {
	var err error
	pgConString := fmt.Sprintf("port=%s host=%s user=%s "+"password=%s dbname=%s sslmode=disable", dbPort, dbHost, dbUsername, dbPassword, dbName)
	db, err = sql.Open("postgres", pgConString)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	if err = completedDataBase(); err != nil {
		log.Fatalf("Error Completed Database: %v\n", err)
	}

	println(weatherResult)
	server := echo.New()
	server.GET("/", index)
	server.GET(WeatherService, weatherService)
	server.GET(DayService, dayService)
	//server.GET("/swagger/*", echoSwagger.WrapHandler)
	server.Logger.Fatal(server.Start(":8084"))
}

//completedDataBase information of weather.
func completedDataBase() error {
	var err error

	if err := database.InitializeDb(db); err != nil {
		log.Fatalf("Error initializing database: %v\n", err)
		return err
	}

	if weatherResult, err = database.GetData(db); len(weatherResult) > 0 || err != nil {
		return err
	}

	if weatherResult, err = weather.GetWeather(db); weatherResult != nil || err != nil {
		return err
	}

	return err
}

//Index API.
func index(server echo.Context) (err error) {
	result := `
	<h1>SISTEMA SOLAR - MercadoLibre</h1>
	<p>Puede dirigirse a la documentacion de la API en la siguiente URL:</p>
	<a href="http://localhost:8084/swagger/index.html">Sistema Solar API swagger!</a>
	<p>Para conocer los periodos de Sequia, Lluvia y condiciones optimas, debe realizarse de la siguiente manera:</p>
	<i>http://localhost:8084/weather</i>	
	<p>Para consultar la información de un dia especifico, debe realizarse de la siguiente manera:</p>
	<i>http://localhost:8084/clima?dia=566</i>
	`
	return server.HTML(http.StatusOK, result)
}

//weatherService return results of weather in JSON format.
func weatherService(server echo.Context) (err error) {
	if len(weatherResult) <= 0 || weatherResult == nil {
		if weatherResult, err = database.GetData(db); err != nil {
			return server.String(http.StatusInternalServerError, "Error: Failed to get weather information from database.")
		}
	}

	return server.JSONPretty(http.StatusOK, weatherResult, " ")
}

//dayService return the result of the weather for a specific day.
func dayService(server echo.Context) (err error) {
	dayParam := server.QueryParam("dia")
	day, err := strconv.Atoi(dayParam)
	if err != nil || day < 0 {
		return server.String(http.StatusBadRequest, "Error: The parameter 'dia' must be a valid number.")
	}

	if day > 3599 {
		return server.String(http.StatusBadRequest, "Error: The parameter 'dia' must be greater than 0 and less than 3599.")
	}

	dayResult := weather.DayResult{}
	if dayResult, err = weather.GetDay(db, day); err != nil {
		return server.String(http.StatusInternalServerError, "Error: Failed to get day information from database.")
	}

	return server.JSON(http.StatusOK, dayResult)
}
