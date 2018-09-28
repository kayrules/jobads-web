package model

import (
	"github.com/globalsign/mgo/bson"
)

type (
	// PricingRules struct
	PricingRules struct {
		ID            bson.ObjectId `json:"id,omitempty" form:"id,omitempty" bson:"_id,omitempty"`
		CustomerID    string        `json:"customer_id,omitempty" form:"customer_id,omitempty" bson:"customer_id,omitempty" valid:"required"`
		ProductCode   string        `json:"product_code,omitempty" form:"product_code,omitempty" bson:"product_code,omitempty" valid:"required"`
		Type          string        `json:"type,omitempty" form:"type,omitempty" bson:"type,omitempty" valid:"required,in(deal|discount)"`
		DealBuy       int           `json:"deal_buy,omitempty" form:"deal_buy,omitempty" bson:"deal_buy,omitempty" valid:"-"`
		DealPriceOf   int           `json:"deal_priceof,omitempty" form:"deal_priceof,omitempty" bson:"deal_priceof,omitempty" valid:"-"`
		DiscountBuy   int           `json:"discount_buy,omitempty" form:"discount_buy,omitempty" bson:"discount_buy,omitempty" valid:"-"`
		DiscountPrice int           `json:"discount_price,omitempty" form:"discount_price,omitempty" bson:"discount_price,omitempty" valid:"-"`
	}
)
