package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type PremiumFeaturesTable_20240205_084704 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &PremiumFeaturesTable_20240205_084704{}
	m.Created = "20240205_084704"

	migration.Register("PremiumFeaturesTable_20240205_084704", m)
}

// Run the migrations
func (m *PremiumFeaturesTable_20240205_084704) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TABLE premium_features (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) default null,
    description VARCHAR(255) default null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME default null
	)`)
}

// Reverse the migrations
func (m *PremiumFeaturesTable_20240205_084704) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE premium_features")
}
