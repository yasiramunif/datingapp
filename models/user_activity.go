package models

// import (
// 	"errors"
// 	"strconv"
// 	"time"
// )

type UserActivityRequest struct {
	OtherUserId  int64 `json:"otherUserId"`
	ActivityType string `json:"activityType"`
}

type UserActivity struct {
	Id           int64  `orm:"pk;auto;column(id)"`
	User         *User  `orm:"column(user_id);rel(one)"`
	OtherUser    *User  `orm:"column(other_user_id);rel(one)"`
	ActivityType string `orm:"column(activity_type);null"`
}

func (u *UserActivity) TableName() string {
	return "user_activities"
}
