package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	var planetMuliplier float64

	switch planet {
	case "Mercury":
		planetMuliplier = 0.2408467
	case "Venus":
		planetMuliplier = 0.61519726
	case "Mars":
		planetMuliplier = 1.8808158
	case "Jupiter":
		planetMuliplier = 11.862615
	case "Saturn":
		planetMuliplier = 29.447498
	case "Uranus":
		planetMuliplier = 84.016846
	case "Neptune":
		planetMuliplier = 164.79132
	default:
		planetMuliplier = 1.0
	}

	const EarthYearInSeconds = 31557600
	return seconds / (EarthYearInSeconds * planetMuliplier)
}
