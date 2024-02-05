package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type UserSubscriptionsTable_20240205_085021 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UserSubscriptionsTable_20240205_085021{}
	m.Created = "20240205_085021"

	migration.Register("UserSubscriptionsTable_20240205_085021", m)
}

// Run the migrations
func (m *UserSubscriptionsTable_20240205_085021) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TABLE user_subscriptions (
		id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT,
		feature_id INT,
		purchased_at DATETIME,		
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		deleted_at DATETIME default null
	)`)
}

// Reverse the migrations
func (m *UserSubscriptionsTable_20240205_085021) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE user_subscriptions")
}
