package model

import "google.golang.org/genproto/googleapis/type/date"

type Order struct {
	userId      int
	orderId     int
	noOfItems   int
	totalAmount int64
	orderDate   date.Date
}
