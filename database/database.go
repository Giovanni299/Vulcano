package database

import "database/sql"

func initializeDb(db *sql.DB) error {
	const weather = `
		CREATE TABLE IF NOT EXISTS "WEATHER" (
			"IdWeather" int,
			"Weather" varchar(10),
			PRIMARY KEY ("IdWeather")
	  	);`

	const days = `
		CREATE TABLE IF NOT EXISTS "DAYS" (
			"IdDay" int,
			"IdWeather" int,
			PRIMARY KEY ("IdDay"),
			CONSTRAINT "fk_Weather"
			FOREIGN KEY("IdWeather") 
	  		REFERENCES "WEATHER"("IdWeather")
	  	);`

	const coordinates = `
		CREATE TABLE IF NOT EXISTS "COORDINATES" (
			"IdCoordinate" int,
			"IdDay" int,
			"xF" float4,
			"yF" float4,
			"xB" float4,
			"yB" float4,
			"xV" float4,
			"yV" float4,
			PRIMARY KEY ("IdCoordinate"),
			CONSTRAINT "fk_Days"
			FOREIGN KEY("IdDay") 
	  		REFERENCES "DAYS"("IdDay")
	  );`

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

func seedDb(db *sql.DB) error {
	const insertWeather = "INSERT INTO days(day, weather) VALUES($1, $2)"
	const insertPlanet = "INSERT INTO planets(name) VALUES($1) RETURNING id"
	const insertCoordinates = "INSERT INTO coordinates(day_id, planet_id, x, y) VALUES ($1, $2, $3, $4)"

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	return tx.Commit()
}
