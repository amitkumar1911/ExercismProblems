//Package weather forecast the current weather condition of various cities in Goblinocus.
package weather


//CurrentCondition denotes the Current weather condition of various  cities in goblinocus.
var CurrentCondition string
// CurrentLocation denotes the Current Location.
var CurrentLocation string


//Forecast function takes two parametres, city and condition and returns a string of current location and current condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
