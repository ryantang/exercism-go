// Package weather provides tools to forecast weather
// conditions by city.
package weather

// CurrentConditions specifies the current conditions as a string.
var CurrentCondition string

// CurrentLocation specifies the city.
var CurrentLocation string

// Forecast takes two parameters: a city and the condition
// It returns a string describing the current weather condition for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
