package weather

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestGetCoordinatesbyDay validated coordinates by day.
func TestGetCoordinatesbyDay(t *testing.T) {
	coordinates := GetCoordinates(0)
	assert.Equal(t, 500.0, coordinates.coorFerengi.x)
	assert.Equal(t, 0.0, coordinates.coorFerengi.y)
	assert.Equal(t, 2000.0, coordinates.coorBetasoide.x)
	assert.Equal(t, 0.0, coordinates.coorBetasoide.y)

	coordinates = GetCoordinates(495)
	assert.Equal(t, -353.553, coordinates.coorFerengi.x)
	assert.Equal(t, -353.553, coordinates.coorFerengi.y)
	assert.Equal(t, 707.107, coordinates.coorVulcano.x)
	assert.Equal(t, -707.107, coordinates.coorVulcano.y)

	coordinates = GetCoordinates(1050)
	assert.Equal(t, 0.0, coordinates.coorBetasoide.x)
	assert.Equal(t, 2000.0, coordinates.coorBetasoide.y)
	assert.Equal(t, -866.025, coordinates.coorVulcano.x)
	assert.Equal(t, -500.0, coordinates.coorVulcano.y)
}
