package main

import (
	"time"
)

/*
func main(){
	// Current epoch time
	fmt.Printf("Current epoch time is:\t\t\t%d\n\n", currentEpochTime())

	// Convert from human readable date to epoch
	humanReadable := time.Now()
	fmt.Printf("Human readable time is:\t\t\t%s\n", humanReadable)
	fmt.Printf("Human readable to epoch time is:\t%d\n\n", humanReadableToEpoch(humanReadable))


	// Convert from epoch to human readable date
	epoch := currentEpochTime()
	fmt.Printf("Epoch to human readable time is:\t%s\n", epochToHumanReadable(epoch))

}
*/

func currentEpochTime() int64 {
	return time.Now().Unix()
}

func humanReadableToEpoch(date time.Time) int64 {
	return date.Unix()
}

func epochToHumanReadable(epoch int64) time.Time {
	return time.Unix(epoch, 0)
}

func humanReadableToEpoch24(day time.Time) int64 {
	return day.Truncate(24 * time.Hour).Unix()
}
