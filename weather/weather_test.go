package weather

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestGetWeather the main function.
func TestGetWeather(t *testing.T) {
	GetWeather(nil)
}

//TestSunInTriangle validate if the planets shape a triangle and the sun is inside.
func TestSunInTriangle(t *testing.T) {
	coordinates := CoordinatesPlanet{
		coorBetasoide: Coordinates{x: 0, y: 2},
		coorFerengi:   Coordinates{x: -2, y: 0},
		coorVulcano:   Coordinates{x: 2, y: 0},
	}
	weather := sunInTriangle(coordinates)
	assert.Equal(t, lluvia, weather)

	coordinates = CoordinatesPlanet{
		coorBetasoide: Coordinates{x: 2, y: 2},
		coorFerengi:   Coordinates{x: 4, y: 2},
		coorVulcano:   Coordinates{x: 2, y: 4},
	}
	weather = sunInTriangle(coordinates)
	assert.Equal(t, ninguno, weather)

	coordinates = CoordinatesPlanet{
		coorBetasoide: Coordinates{x: -0.5, y: 0},
		coorFerengi:   Coordinates{x: 0, y: 0.5},
		coorVulcano:   Coordinates{x: 0.5, y: -0.5},
	}
	weather = sunInTriangle(coordinates)
	assert.Equal(t, lluvia, weather)

	coordinates = CoordinatesPlanet{
		coorBetasoide: Coordinates{x: 0, y: 2},
		coorFerengi:   Coordinates{x: 0, y: -2},
		coorVulcano:   Coordinates{x: 2, y: 0},
	}
	weather = sunInTriangle(coordinates)
	assert.Equal(t, lluvia, weather)

	coordinates = CoordinatesPlanet{
		coorBetasoide: Coordinates{x: -2000, y: 0},
		coorFerengi:   Coordinates{x: 2000, y: 0},
		coorVulcano:   Coordinates{x: 0, y: 600},
	}
	weather = sunInTriangle(coordinates)
	assert.Equal(t, lluviaIntensa, weather)
}

//TestPlanetsAligned validated if the planets are aligned.
func TestPlanetsAligned(t *testing.T) {
	coordinates := CoordinatesPlanet{
		coorBetasoide: Coordinates{x: 2000, y: 0},
		coorFerengi:   Coordinates{x: 500, y: 0},
		coorVulcano:   Coordinates{x: 1000, y: 0},
	}
	_, weather := planetsAligned(coordinates)
	assert.Equal(t, sequia, weather)

	coordinates = CoordinatesPlanet{
		coorBetasoide: Coordinates{x: -1956.295, y: 415.823},
		coorFerengi:   Coordinates{x: 219.186, y: -449.397},
		coorVulcano:   Coordinates{x: 766.044, y: -642.788},
	}
	_, weather = planetsAligned(coordinates)
	assert.Equal(t, optimo, weather)

	coordinates = CoordinatesPlanet{
		coorBetasoide: Coordinates{x: 1414.214, y: -1414.214},
		coorFerengi:   Coordinates{x: 353.553, y: -353.553},
		coorVulcano:   Coordinates{x: -707.107, y: 707.107},
	}
	_, weather = planetsAligned(coordinates)
	assert.Equal(t, sequia, weather)

	coordinates = CoordinatesPlanet{
		coorBetasoide: Coordinates{x: 2, y: 0},
		coorFerengi:   Coordinates{x: 2, y: 2},
		coorVulcano:   Coordinates{x: 2, y: 4},
	}
	_, weather = planetsAligned(coordinates)
	assert.Equal(t, optimo, weather)

	coordinates = CoordinatesPlanet{
		coorBetasoide: Coordinates{x: 0, y: -2},
		coorFerengi:   Coordinates{x: 2, y: 0},
		coorVulcano:   Coordinates{x: 4, y: 2},
	}
	_, weather = planetsAligned(coordinates)
	assert.Equal(t, optimo, weather)
}
