package models

// TableName function to define User table name
func (u *User) TableName() string {
	return "users"
}

type User struct {
	Id       int64   `orm:"pk;auto;column(id)"`
	Username string  `orm:"column(username);null"`
	Email    string  `orm:"column(email);null"`
	Password string  `json:"-" orm:"column(password);null"`
	Profile  *Profile `orm:"reverse(one);null"`
}

type Profile struct {
	ProfileId int64  `orm:"pk;auto;column(id)"`
	Gender    string `orm:"column(gender);null"`
	BirthDate string `orm:"column(birthdate);null"`
	ImgPath   string `orm:"column(img_path);null"`
	Fullname  string `orm:"column(fullname);null"` 
	User      *User  `orm:"column(user_id);rel(one)"`
}

func (u *Profile) TableName() string {
	return "user_profiles"
}
