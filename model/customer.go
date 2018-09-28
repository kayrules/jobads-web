package model

type (
	// Customer struct
	Customer struct {
		ID   string `json:"id" bson:"id"`
		Name string `json:"name" bson:"name" valid:"required"`
	}
)
