// Package weather is made for showing weather forecast in some location.
package weather

// CurrentCondition holds current weather forecast condition.
var CurrentCondition string

// CurrentLocation holds current weather forecast location.
var CurrentLocation string

// Forecast shows current forecast based on location and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
