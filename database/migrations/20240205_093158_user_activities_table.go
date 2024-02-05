package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type UserActivitiesTable_20240205_093158 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UserActivitiesTable_20240205_093158{}
	m.Created = "20240205_093158"

	migration.Register("UserActivitiesTable_20240205_093158", m)
}

// Run the migrations
func (m *UserActivitiesTable_20240205_093158) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TABLE user_activities (
		id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT,
		activity_type ENUM('Pass', 'Like'),
		other_user_id INT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		deleted_at DATETIME default null
	)`)
}

// Reverse the migrations
func (m *UserActivitiesTable_20240205_093158) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE user_activities")
}
