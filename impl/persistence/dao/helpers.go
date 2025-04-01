package dao

import (
	"math"
	"strconv"

	log "github.com/sirupsen/logrus"
)

const (
	Sigma      = 5.670e-8 // Stefan-Boltzmann constant (W/m²K⁴)
	Lsun       = 3.828e26 // Solar luminosity (W)
	RsunMeters = 6.955e8  // Solar radius (m)
	UpperConst = 1.1      // Solar flux scaling factor for upper bound of habitable zone
	LowerConst = 0.53     // Solar flux scaling factor for lower bound of habitable zone
)

// Gets the string value of a CSV column
func GetStringValue(record []string, colIndices map[string]int, colName string) string {
	if idx, exists := colIndices[colName]; exists && idx < len(record) {
		return record[idx]
	}
	return ""
}

// Gets the float value of a CSV column
func GetFloatValue(record []string, colIndices map[string]int, colName string) float32 {
	if idx, exists := colIndices[colName]; exists && idx < len(record) {
		return ParseFloat(record[idx])
	}
	return 0.0
}

// Parse float from strings
func ParseFloat(val string) float32 {
	if val == "" {
		return 0.0
	}

	floatVal, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return 0.0
	}

	return float32(floatVal)
}

// Get habitable zone of star
func (s *Star) GetHabitableZone() (lower float64, upper float64) {
	radius, temp := float64(s.Radius*RsunMeters), float64(s.Temp)
	luminosity := (4 * math.Pi * math.Pow(radius, 2) * Sigma * math.Pow(temp, 4)) / Lsun

	lowerBound, upperBound := math.Sqrt(luminosity/1.1), math.Sqrt(luminosity/0.53)

	log.Debugf("Habitable zone of star %v with luminosity %f: [%f, %f]", s.ID, luminosity, lowerBound, upperBound)
	return lowerBound, upperBound
}

// Check if exoplanet is in the habitable zone of its star
func (e *Exoplanet) IsInHabitableZone(s *Star) bool {
	lower, upper := s.GetHabitableZone()
	return e.Dist >= float32(lower) && e.Dist <= float32(upper)
}
