package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SubscriptionType struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Icon        string             `bson:"icon" json:"icon"`
	Name        string             `bson:"name" json:"name"`
	Payments    []PaymentType      `bson:"payments" json:"payments"`
	CreatedTime time.Time          `bson:"created_time" json:"created_time"`
}

type PaymentType struct {
	PlanName        string  `bson:"plan_name" json:"plan_name"`
	Price           float64 `bson:"price" json:"price"`
	DiscountedPrice float64 `bson:"discounted_price" json:"discounted_price"`
}
