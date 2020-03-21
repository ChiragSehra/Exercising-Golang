package space

// Planet represents the planet on which we want to calculate the age
type Planet string

// Age calcylates the age relative to planet rotation
func Age(ageInSeconds float64, planet Planet) float64 {
	const earthRevolutions = 31557600

	// creating a dictionary of planets and there age

	Planets := map[Planet]float64{
		"Earth":   earthRevolutions,
		"Mercury": earthRevolutions * 0.2408467,
		"Venus":   earthRevolutions * 0.61519726,
		"Mars":    earthRevolutions * 1.8808158,
		"Jupiter": earthRevolutions * 11.862615,
		"Saturn":  earthRevolutions * 29.447498,
		"Uranus":  earthRevolutions * 84.016846,
		"Neptune": earthRevolutions * 164.79132,
	}

	return calculateAge(Planets[planet], ageInSeconds)
}

func calculateAge(planetRevolution float64, ageInSeconds float64) float64 {
	return ageInSeconds / planetRevolution
}
