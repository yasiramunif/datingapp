package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type UserTable_20240203_213450 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UserTable_20240203_213450{}
	m.Created = "20240203_213450"

	migration.Register("UserTable_20240203_213450", m)
}

// Run the migrations
func (m *UserTable_20240203_213450) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TABLE users (
		id INT AUTO_INCREMENT PRIMARY KEY, 
		username VARCHAR(255) default null, 
		password VARCHAR(255) default null, 
		email VARCHAR(255) default null,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		deleted_at DATETIME default null,
		unique (username, email)
	)`)
}

// Reverse the migrations
func (m *UserTable_20240203_213450) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE users")
}
