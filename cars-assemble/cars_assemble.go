package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// panic("CalculateWorkingCarsPerHour not implemented")

    var ans float64;

    var floatProductionRate=float64(productionRate)

    ans=(successRate*floatProductionRate)/100;

    return ans;
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	// panic("CalculateWorkingCarsPerMinute not implemented")

    var ans float64;

    ans=CalculateWorkingCarsPerHour(productionRate,successRate )/60;

    var intResult=int(ans);

    return intResult;
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// panic("CalculateCost not implemented")

    var ans uint

    var cars uint ;
    cars=uint(carsCount)

    cars=cars/10

    var remCars uint
    remCars=uint(carsCount)-(cars*10)

    ans=cars*95000+remCars*10000;

    return ans;
}
