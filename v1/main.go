package main

import (
	"booking-client-grpc-api/v1/client"
	"encoding/json"
	"fmt"
)

func main() {
	c := client.CreateClient()
	bookingDetails := client.MakeBooking(c)
	bookingBytes, _ := json.Marshal(bookingDetails)
	fmt.Println(string(bookingBytes))

	fmt.Println("Fetching bookings by location: ", "Bengaluru, India")
	bookingsByLocation := client.GetAllBookingsByLocation(c, "Bengaluru, India")
	bookingByLocationBytes, _ := json.Marshal(bookingsByLocation)
	fmt.Println(string(bookingByLocationBytes))

	fmt.Println("Fetching bookings by booking name: ", "Shyam")
	bookingsByName := client.GetAllBookingsByName(c, "Shyam")
	bookingByNameBytes, _ := json.Marshal(bookingsByName)
	fmt.Println(string(bookingByNameBytes))

	fmt.Println("Fetching bookings by movie: ", "top gun: maverick")
	bookingsByMovie := client.GetAllBookingsByMovie(c, "top gun: maverick")
	bookingByMovieBytes, _ := json.Marshal(bookingsByMovie)
	fmt.Println(string(bookingByMovieBytes))

}
