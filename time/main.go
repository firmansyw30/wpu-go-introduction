package main

import (
	"fmt"
	"time"
)

func getUTCCurrentTime() {
	fmt.Println("Sample Time in Go")
	now := time.Now().UTC()
	fmt.Println("Current Time in UTC: ", now.String())
}

func getCurrentTime() {
	now := time.Now()                           // Get the current local time
	fmt.Println("Current Time: ", now.String()) // Print the current time as a string

	// Get the current time in a specific timezone (e.g., "Asia/Jakarta")
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Error loading location: ", err.Error())
		return
	}
	nowInJakarta := now.In(location)                                     // Convert the current time to the specified timezone
	fmt.Println("Current Time in Asia/Jakarta: ", nowInJakarta.String()) // Print the current time in the specified timezone as a string

}

func convertToZone(t time.Time, zone string) time.Time {
	loc, err := time.LoadLocation(zone)
	if err != nil {
		fmt.Println("Error in loading location: ", err)
		return t
	}
	return t.In(loc)
}

func zoneConversion() {
	println("Sample Time Zone Conversion in Go")
	nowUTC := time.Now().UTC()
	fmt.Println("Current UTC time: ", nowUTC)
	// Convert UTC time to Asia/Jakarta time
	jakartaTime := convertToZone(nowUTC, "Asia/Jakarta")
	fmt.Println("Current time in Asia/Jakarta: ", jakartaTime)
	// Convert UTC time to America/New_York time (DST, Daylight Saving Time)
	newYorkTime := convertToZone(nowUTC, "America/New_York")
	fmt.Println("Current time in America/New_York: ", newYorkTime)
	//Store time (simulation store to Database)
	fmt.Println("Stored to Database (UTC): ", nowUTC)
}

func main() {
	getCurrentTime()
	fmt.Println()
	getUTCCurrentTime()
	fmt.Println()
	zoneConversion()
}
