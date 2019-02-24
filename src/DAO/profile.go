package DAO

import (
	"fmt"

	"../connection"
	"gopkg.in/mgo.v2/bson"
)

const (
	collection = "Profile"
)

// ProfileDAO is an exported type that
// contains ID, Name and Phone
type ProfileDAO struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name  string        `json:"name" bson:"name"`
	Phone string        `json:"phone" bson:"phone"`
}

// Create method that create a Profile
func (p *ProfileDAO) Create() error {
	err := connection.Db.C(collection).Insert(p)
	if err != nil {
		return fmt.Errorf("Error to create Profile: %v", err)
	}
	return err
}
