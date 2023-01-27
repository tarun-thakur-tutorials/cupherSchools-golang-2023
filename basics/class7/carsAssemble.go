package main

/*
In this exercise you'll be writing code to analyze the production of an assembly line in a car factory. The assembly line's speed can range from 0 (off) to 10 (maximum).

At its default speed (1), 221 cars are produced each hour. In principle, the production increases linearly. So with the speed set to 4, it should produce 4 * 221 = 884 cars per hour. However, higher speeds increase the likelihood that faulty cars are produced, which then have to be discarded.

Also, there are times when the assembly line has an artificially imposed limit on the throughput (meaning no more than the limit can be produced per hour).

1. Calculate the success rate
Implement a function (SuccessRate) to calculate the ratio of an item being created without error for a given speed. The following table shows how speed influences the success rate:

0: 0% success rate.
1 - 4: 100% success rate.
5 - 8: 90% success rate.
9 - 10: 77% success rate.
rate := SuccessRate(6)
fmt.Println(rate)
// Output: 0.9
2. Calculate the production rate per hour
Implement a function to calculate the assembly line's production rate per hour:

rate := CalculateProductionRatePerHour(7)
fmt.Println(rate)
// Output: 1392.3
Note that the value returned is of type float64.

3. Calculate the number of working items produced per minute
Implement a function to calculate how many cars are produced each minute:

rate := CalculateProductionRatePerMinute(5)
fmt.Println(rate)
// Output: 16
Note that the value returned is of type int.

4. Calculate the artificially-limited production rate
Implement a function to calculate the assembly line's production rate per hour:

rate := CalculateLimitedProductionRatePerHour(2, 1000.0)
fmt.Println(rate)
// Output: 442.0
rate := CalculateLimitedProductionRatePerHour(7, 1000.0)
fmt.Println(rate)
// Output: 1000.0
*/
const CarPerHour = 221

func SuccessRate(speed int) (percentage float64) {

	if speed <= 0 {
		percentage = 0
		return
	}

	if speed >= 1 && speed <= 4 {
		percentage = 1 // 100 %
	} else if speed <= 8 {
		percentage = 0.9 // 90%
	} else if speed <= 10 {
		percentage = 0.77
	}

	return
}

func SuccessRate2(speed int) float64 {
	switch speed {
	case 1, 2, 3, 4:
		return 1
	case 5, 6, 7, 8:
		return 0.9
	case 9, 10:
		return 0.77
	default:
		return 0
	}

}

func CalculateProductionRatePerHour(speed int) float64 {
	return float64(CarPerHour*speed) * SuccessRate(speed)
}
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / 60)
}

func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	production := CalculateProductionRatePerHour(speed)

	if production > limit {
		return limit
	} else {
		return production
	}
}
