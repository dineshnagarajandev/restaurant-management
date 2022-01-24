package models

import (
	"time"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID             primitive.ObjectID `bson:"_id"`
	Invoice_id     string             `json:"invoice_id"`
	Order_id       string             `json:"order_id"`
	Payment_method *string            `json:"payment_method" validate:"eq=CASH|eq=CARD|eq="`
	Payment_status *string            `json:"payment_status"`
	Payment_date   time.Time          `json:"payment_date"`
	Created_at     time.Time          `json:"created_at"`
	Updated_at     time.Time          `json:"updated_at"`
}
