package models

// import (
// 	"errors"
// 	"strconv"
// 	"time"
// )

type SubscriptionRequest struct {
	FeatureId int64 `json:"featureId"`
}

type Subscription struct {
	Id          int64   `orm:"pk;auto;column(id)"`
	User        *User  `orm:"column(user_id);rel(one)"`
	Feature     *Feature `orm:"column(feature_id);rel(one)"`
	PurchasedAt string  `orm:"column(purchased_at);null"`
}

func (u *Subscription) TableName() string {
	return "user_subscriptions"
}
