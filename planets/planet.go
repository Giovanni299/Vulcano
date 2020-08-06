package planet

const (
	//FerengiDaysYear is the number of days required to complete a solar year on the planet Ferengis.
	FerengiDaysYear = 3600

	//MaxPerimeterTriangle in km, the maximum value is 6262.30068km but I leave this value to give a greater tolerance of LluviaIntensa.
	MaxPerimeterTriangle = 6262
)

//Planet struct information planet.
type Planet struct {
	Name            string
	AngularSpeed    uint
	OrbitRadius     float64
	HoraryDirection bool
}

//WeatherResult data struct of table weather of DB.
type WeatherResult struct {
	Weather string
	Value   string
}

var (
	//Planets information in the solar system.
	Planets = []Planet{
		{
			Name:            "Ferengi",
			AngularSpeed:    1,
			OrbitRadius:     500,
			HoraryDirection: true,
		},
		{
			Name:            "Betasoide",
			AngularSpeed:    3,
			OrbitRadius:     2000,
			HoraryDirection: true,
		},
		{
			Name:            "Vulcano",
			AngularSpeed:    5,
			OrbitRadius:     1000,
			HoraryDirection: false,
		},
	}
)
