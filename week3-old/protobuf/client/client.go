package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "week3/CustomerFeedback"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

var c pb.CustomerFeedbackClient

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c = pb.NewCustomerFeedbackClient(conn)
	// r, err := c.GetFeedbackByPassagerID(ctx, &pb.GetFeedbackByPassagerIDRequest{PassengerID: PassagerID})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Printf("GetFeedbackByPassagerID  response: %s", r.Feedbacks)
	//reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("----Choose Menu----")
		fmt.Println("1 : Add Feedback")
		fmt.Println("2 : Get Feedback By BookingCode")
		fmt.Println("3 : Get Feedback By PassengerID")
		fmt.Println("4 : Delete Feedback")
		var menuID int
		_, err := fmt.Scan(&menuID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("You choosed : ", menuID)
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		var passengerID int32
		var bookingCode string
		var feedback string
		switch menuID {
		case 1:
			fmt.Println("-- Please input PassengerID--")
			fmt.Scan(&passengerID)
			fmt.Println("-- Please input BookingCode--")
			fmt.Scan(&bookingCode)
			fmt.Println("-- Please input Feedback--")
			fmt.Scan(&feedback)
			addFeedback(ctx, &pb.PassengerFeedback{
				PassengerID: passengerID,
				BookingCode: bookingCode,
				Feedback:    feedback})
		case 2:
			fmt.Println("-- Please input BookingCode--")
			fmt.Scan(&bookingCode)
			getFeedbackByBookingCode(ctx, bookingCode)
		case 3:
			var passengerID int32
			fmt.Println("-- Please input PassengerID--")
			fmt.Scan(&passengerID)
			getFeedbackByPassagerID(ctx, passengerID)
		}
		break
	}
}
func getFeedbackByBookingCode(ctx context.Context, bookingCode string) {
	r, err := c.GetFeedbackByBookingCode(ctx, &pb.GetFeedbackByBookingCodeRequest{BookingCode: bookingCode})
	if err != nil {
		log.Fatalf("could not get feed back by passager id: %v", err)
	}
	log.Printf("GetFeedbackByBookingCode  response: %s", r.GetPassengerFeedback())
}
func getFeedbackByPassagerID(ctx context.Context, passengerID int32) {
	r, err := c.GetFeedbackByPassengerID(ctx, &pb.GetFeedbackByPassengerIDRequest{PassengerID: passengerID})
	if err != nil {
		log.Fatalf("could not get feed back by passager id: %v", err)
	}
	log.Printf("GetFeedbackByPassagerID  response: %s", r.Feedbacks)
}
func addFeedback(ctx context.Context, passengerFeedback *pb.PassengerFeedback) {
	r, err := c.AddFeedback(ctx, &pb.AddFeedbackRequest{PassengerFeedback: passengerFeedback})
	if err != nil {
		log.Fatalf("could not get feed back by passager id: %v", err)
	}
	log.Printf("GetFeedbackByPassagerID  response: %s", r.GetMsg())
}
