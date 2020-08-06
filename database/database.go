package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	planet "github.com/Giovanni299/Vulcano/planets"
)

const (
	weather = `
		CREATE TABLE IF NOT EXISTS weather(
			idWeather int, weather varchar(15), value varchar(250),
			PRIMARY KEY (idWeather)
		  );`

	days = `
		  CREATE TABLE IF NOT EXISTS days (
			  idDay int, idWeather int, 
			  PRIMARY KEY (idDay),
			  CONSTRAINT "fk_Weather"
			  FOREIGN KEY(idWeather) 
				REFERENCES weather(idWeather)
			);`

	coordinates = `
		CREATE TABLE IF NOT EXISTS coordinates (
			idCoordinate serial primary key, idDay int, xF float4, yF float4,
			xB float4, yB float4, xV float4, yV float4,
			CONSTRAINT "fk_Days"
			FOREIGN KEY(idDay) 
	  		REFERENCES days(idDay)
	  );`

	weatherType = `INSERT INTO weather (idWeather , weather, value) 
	VALUES
		(0, 'Sequia', $1),
		(1, 'Lluvia', $2),
		(2, 'Lluvia Intensa', $3),
		(3, 'Optimo', $4),
		(4, 'Ninguno', '--')  ON CONFLICT DO NOTHING;`

	insertWeather     = `INSERT INTO days(idDay, idWeather) VALUES%s`
	insertCoordinates = `INSERT INTO coordinates(idday, xb, yb, xf, yf, xv, yv) VALUES%s`
	selectWeather     = `SELECT weather, value FROM weather LIMIT $1`
	selectDay         = `SELECT idweather from days where idday = $1;`
)

//InitializeDb tables on DataBase.
func InitializeDb(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	queries := []string{weather, days, coordinates}

	for _, q := range queries {
		_, err = tx.Exec(q)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

//GetData from DB.
func GetData(db *sql.DB) ([]planet.WeatherResult, error) {
	weatherResults := []planet.WeatherResult{}
	rows, err := db.Query(selectWeather, 4)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var weather string
	var value string
	for rows.Next() {
		err = rows.Scan(&weather, &value)
		if err != nil {
			return nil, err
		}

		weatherResults = append(weatherResults, planet.WeatherResult{Weather: weather, Value: value})
	}

	return weatherResults, err
}

//GetDay return weather od DB for a specific day.
func GetDay(db *sql.DB, day int) (int, error) {
	idweather := 0
	err := db.QueryRow(selectDay, day).Scan(&idweather)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	return idweather, err
}

//InsertData in DB.
func InsertData(db *sql.DB, value []string, arguments []interface{}, valueCoordinates []string, valueCoorArgs []interface{}, resultWeatherArgs []interface{}) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(weatherType, resultWeatherArgs...)
	if err != nil {
		tx.Rollback()
		return err
	}

	stmt := fmt.Sprintf(insertWeather, strings.Join(value, ","))
	_, err = tx.Exec(stmt, arguments...)
	if err != nil {
		tx.Rollback()
		return err
	}

	stmt = fmt.Sprintf(insertCoordinates, strings.Join(valueCoordinates, ","))
	_, err = tx.Exec(stmt, valueCoorArgs...)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
