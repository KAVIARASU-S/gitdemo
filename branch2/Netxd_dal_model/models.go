package netxddalmodel

import (

	"time"
	//"go.monogodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	Customer_ID int   `json:"customerid"`
	First_Name string  `json:"firstname"`
	Last_Name string    `json:"lastname"`
	Bank_ID int          `json:"bankid"`
	Balance int  `json:"balance"`
	Credited_At time.Time `json:"creditedat"`
    Updated_At time.Time   `json:"updatedat"`
	IS_Active bool      `json:"isactive"`
}

type CustomerResponse struct{
	CustomerID int  `json:"customerid"`
	CreatedAt  string  `json:"createdat"`
}