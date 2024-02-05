package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type UserProfilesTable_20240204_191054 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UserProfilesTable_20240204_191054{}
	m.Created = "20240204_191054"

	migration.Register("UserProfilesTable_20240204_191054", m)
}

// Run the migrations
func (m *UserProfilesTable_20240204_191054) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TABLE user_profiles (
		id INT AUTO_INCREMENT PRIMARY KEY, 
		img_path VARCHAR(255) default null,
		gender VARCHAR(255) default null,
		fullname VARCHAR(255) default null,
		birthdate DATE default null,
		user_id INT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		deleted_at DATETIME default null,
		CONSTRAINT FK_UsersUserProfiles FOREIGN KEY (user_id)
    REFERENCES users(id)
	)`)
}

// Reverse the migrations
func (m *UserProfilesTable_20240204_191054) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE user_profiles")
}
