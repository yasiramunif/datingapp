package models

// import (
// 	"errors"
// 	"strconv"
// 	"time"
// )

type Feature struct {
	Id          int64   `orm:"pk;auto;column(id)"`
	Name        string  `orm:"column(name);null"`
	Description string  `orm:"column(description);null"`
}

func (u *Feature) TableName() string {
	return "premium_features"
}
