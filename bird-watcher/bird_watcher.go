package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	// panic("Please implement the TotalBirdCount() function") 

    count:=0

    for i:=0;i<len(birdsPerDay);i++{


        count+=birdsPerDay[i]
        
    }

    return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	// panic("Please implement the BirdsInWeek() function") 

    lastDay:=7*week 

    sum:=0
    start:=0 

    if week!=1{
        start=lastDay-7
    }

    for  i:=start;i<lastDay;i++{
        sum+=birdsPerDay[i]
    }

    return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	// panic("Please implement the FixBirdCountLog() function")

    day:=0 
    for i:=0;i<len(birdsPerDay);i++{
        if i==day{
            birdsPerDay[i]+=1 
            day+=2
        }
    }

    return birdsPerDay
}
