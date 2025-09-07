package models

// TripType نوع سفر

type TripType string

const (
	TripTypeAirplane TripType = "airplane"
	TripTypeTrain    TripType = "train"
	TripTypeBus      TripType = "bus"
)

// Trip مدل سفر
