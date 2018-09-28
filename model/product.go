package model

import (
	"github.com/globalsign/mgo/bson"
)

type (
	// Product struct
	Product struct {
		ID    bson.ObjectId `json:"id,omitempty" form:"id,omitempty" bson:"_id,omitempty"`
		Code  string        `json:"code,omitempty" form:"code,omitempty" bson:"code,omitempty" valid:"required"`
		Name  string        `json:"name,omitempty" form:"name,omitempty" bson:"name,omitempty" valid:"required"`
		Price int           `json:"price,omitempty" form:"price,omitempty" bson:"price,omitempty" valid:"int,required"`
	}
)
