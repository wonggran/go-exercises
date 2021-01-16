package space

type Planet string
type OrbitalPeriod float64

// Planet name to # of years for an Earth orbital period.
var conversionTable = map[Planet]OrbitalPeriod{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func (p Planet) OrbitalPeriod() OrbitalPeriod {
	return conversionTable[p]
}

func Age(earthSeconds float64, planet Planet) float64 {
	numSecondsPerYear := float64(60 * 60 * 24 * 365)
	earthSecondsToYears := earthSeconds / numSecondsPerYear
	planetFactor := planet.OrbitalPeriod()
	numYearsOnPlanet := earthSecondsToYears * float64(planetFactor)
	return numYearsOnPlanet
}
