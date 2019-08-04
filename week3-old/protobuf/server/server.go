package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "week3/CustomerFeedback"
	model "week3/protobuf/server/models"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	port = ":50051"
)

type server struct{}

var feedbacks = []pb.PassengerFeedback{
	pb.PassengerFeedback{
		BookingCode: "1",
		PassengerID: 1,
		Feedback:    "Good",
	},
	pb.PassengerFeedback{
		BookingCode: "2",
		PassengerID: 2,
		Feedback:    "Normal",
	},
	pb.PassengerFeedback{
		BookingCode: "3",
		PassengerID: 2,
		Feedback:    "Bad",
	},
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:root@/PassengerFeedback?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal("failed to connect db")
	}
	db.LogMode(true)
	err = db.AutoMigrate(model.PassengerFeedback{}).Error
	if err != nil {
		log.Fatal("failed to migrate table todos")
	}
}

func (s *server) AddFeedback(ctx context.Context, in *pb.AddFeedbackRequest) (*pb.AddFeedbackResponse, error) {
	passangerFeedback := model.PassengerFeedback{
		Feedback:    in.PassengerFeedback.GetFeedback(),
		BookingCode: in.PassengerFeedback.GetBookingCode(),
		PassengerID: in.PassengerFeedback.GetPassengerID()}
	var filteredFeedback pb.PassengerFeedback
	if db.Where(&model.PassengerFeedback{BookingCode: in.PassengerFeedback.GetBookingCode()}).Find(&filteredFeedback).RecordNotFound() {
		db.Save(&passangerFeedback)
		return &pb.AddFeedbackResponse{Msg: "Add feedback successfully " + in.PassengerFeedback.GetBookingCode()}, nil
	} else {
		return &pb.AddFeedbackResponse{Msg: "Add feedback failed, already have feedback for bookingcode " + in.PassengerFeedback.GetBookingCode()}, nil
	}
}

func (s *server) GetFeedbackByBookingCode(ctx context.Context, in *pb.GetFeedbackByBookingCodeRequest) (*pb.GetFeedbackByBookingCodeResponse, error) {
	bookingCode := in.GetBookingCode()
	fmt.Println("server received bookingCode: ", bookingCode)
	var filteredFeedback pb.PassengerFeedback
	db.Where(&model.PassengerFeedback{BookingCode: bookingCode}).First(&filteredFeedback)
	return &pb.GetFeedbackByBookingCodeResponse{PassengerFeedback: &filteredFeedback}, nil
}

func (s *server) GetFeedbackByPassengerID(ctx context.Context, in *pb.GetFeedbackByPassengerIDRequest) (*pb.GetFeedbackByPassengerIDResponse, error) {
	passengerID := in.GetPassengerID()
	fmt.Println("server received passengerID: ", passengerID)
	var filteredFeedbacks []*pb.PassengerFeedback
	db.Where(&model.PassengerFeedback{PassengerID: passengerID}).Find(&filteredFeedbacks)
	return &pb.GetFeedbackByPassengerIDResponse{Feedbacks: filteredFeedbacks}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCustomerFeedbackServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
