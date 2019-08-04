package models

import "github.com/jinzhu/gorm"

type PassengerFeedback struct {
	gorm.Model
	PassengerID int32  `json:"passagerID"`
	BookingCode string `json:"bookingCode"`
	Feedback    string `json:"feedback"`
}
