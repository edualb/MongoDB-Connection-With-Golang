package dao

import (
	"fmt"
	"time"
)

const (
	collection = "Profile"
)

// Profile is an exported type that
// contains ID, Name and Phone
type Profile struct {
	Name  string    `json:"name" bson:"name"`
	Phone string    `json:"phone" bson:"phone"`
	Date  time.Time `json:"time" bson:"time"`
}

// Create method that create a Profile
func (p *Profile) Create() error {
	err := Db.C(collection).Insert(p)
	if err != nil {
		return fmt.Errorf("Error to create Profile: %v", err)
	}
	return err
}
