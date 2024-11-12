package models

import "time"

type Payment struct {
	Id int `json:"id"`
	PurchaseId int `json:"purchase_id"`
	Amount int `json:"amount"`
	PaymentMethod string `json:"payment_method"`
	PaymentDate time.Time `json:"payment_date"`
	Status string `json:"status"`
}

type CreatePayment struct {
	PurchaseId int `json:"purchase_id"`
	Amount int `json:"amount"`
	PaymentMethod string `json:"payment_method"`
	PaymentDate time.Time `json:"payment_date"`
	Status string `json:"status"`
}

type UpdatePayment struct {
	Id          int    `json:"id"`
	PurchaseId int `json:"purchase_id"`
	Amount int `json:"amount"`
	PaymentMethod string `json:"payment_method"`
	PaymentDate time.Time `json:"payment_date"`
	Status string `json:"status"`
}

type GetListPaymentRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListPaymentResponse struct {
	Count   int        `json:"count"`
	Payment []*Payment `json:"Payment"`
}
