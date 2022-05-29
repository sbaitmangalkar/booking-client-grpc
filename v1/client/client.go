package client

import (
	pb "booking-client-grpc-api/v1/grpc"
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:8980", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func GetBookingInformation(client pb.MovieBookingServiceClient) pb.BookingDetails {
	ctx := context.Background()
	bookingQuery1 := pb.BookingQuery{}
	bookingQuery1.Id = ""

	bookingDetails1, _ := client.GetBookingDetails(ctx, &bookingQuery1)
	return *bookingDetails1
}

func GetAllBookingsByName(client pb.MovieBookingServiceClient, name string) []pb.BookingDetails {
	res := make([]pb.BookingDetails, 0)
	ctx := context.Background()
	query := pb.BookingQuery{}
	query.BookingName = name
	bookings, err := client.GetAllBookingsByName(ctx, &query)
	if err != nil {
		log.Fatalf("Error occurred while fetching booking for name: %s", query.BookingName)
	}
	if bookings != nil {
		for {
			bookingDetails, receiveError := bookings.Recv()
			if receiveError == io.EOF {
				break
			}
			if receiveError != nil {
				log.Fatalf("bookings.Recv call failed with error %s", receiveError.Error())
			}
			res = append(res, *bookingDetails)
		}
	}
	return res
}

func GetAllBookingsByLocation(client pb.MovieBookingServiceClient, location string) []pb.BookingDetails {
	res := make([]pb.BookingDetails, 0)
	ctx := context.Background()
	query := pb.BookingQuery{}
	query.Location = location
	bookings, err := client.GetAllBookingsByLocation(ctx, &query)
	if err != nil {
		log.Fatalf("Error occurred while fetching booking for location: %s", query.Location)
	}
	if bookings != nil {
		for {
			bookingDetails, receiveError := bookings.Recv()
			if receiveError == io.EOF {
				break
			}
			if receiveError != nil {
				log.Fatalf("bookings.Recv call failed with error %s", receiveError.Error())
			}
			res = append(res, *bookingDetails)
		}
	}
	return res
}

func GetAllBookingsByMovie(client pb.MovieBookingServiceClient, movie string) []pb.BookingDetails {
	res := make([]pb.BookingDetails, 0)
	ctx := context.Background()
	query := pb.BookingQuery{}
	query.Location = movie
	bookings, err := client.GetAllBookingsByMovie(ctx, &query)
	if err != nil {
		log.Fatalf("Error occurred while fetching booking for location: %s", query.Location)
	}
	if bookings != nil {
		for {
			bookingDetails, receiveError := bookings.Recv()
			if receiveError == io.EOF {
				break
			}
			if receiveError != nil {
				log.Fatalf("bookings.Recv call failed with error %s", receiveError.Error())
			}
			res = append(res, *bookingDetails)
		}
	}
	return res
}

func MakeBooking(client pb.MovieBookingServiceClient) []pb.BookingDetails {
	ctx := context.Background()
	res := make([]pb.BookingDetails, 0)
	for i := 1; i <= 10; i++ {
		movie := "Dr. Strange in the Multiverse Of Madness"
		location := "Bengaluru, India"
		date := "02-06-2022"
		if i%2 != 0 {
			movie = "Top Gun: Maverick"
			location = "Bengaluru, India"
			date = "03-06-2022"
		}
		query := pb.BookingQuery{}
		query.BookingEmail = "sbaitman@gmail.com"
		query.BookingName = "Shyam"
		query.Movie = movie
		query.Date = date
		query.Location = location

		bookingDetails, err1 := client.MakeBooking(ctx, &query)
		if err1 != nil {
			log.Println("Error occurred while making grpc call", err1)
		}
		res = append(res, *bookingDetails)
	}
	return res
}

func CreateClient() pb.MovieBookingServiceClient {
	var opts []grpc.DialOption
	if *tls {
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	//defer conn.Close()
	client := pb.NewMovieBookingServiceClient(conn)
	return client
}
