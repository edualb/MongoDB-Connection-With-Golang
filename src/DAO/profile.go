package dao

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	collection = "Profile"
)

// Profile is an exported type that
// contains ID, Name and Phone
type Profile struct {
	ID    bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name  string        `json:"name" bson:"name"`
	Phone string        `json:"phone" bson:"phone"`
	Date  time.Time     `json:"time" bson:"time"`
}

// Create method that create a Profile
func (p *Profile) Create() error {
	err := Db.C(collection).Insert(p)
	if err != nil {
		return fmt.Errorf("Error to create Profile: %v", err)
	}
	return err
}

// GetAll get all Profile documents
func (p *Profile) GetAll() ([]Profile, error) {
	var profiles []Profile
	err := Db.C(collection).Find(bson.M{}).All(&profiles)
	if err != nil {
		return nil, fmt.Errorf("Error to create Profile: %v", err)
	}
	return profiles, err
}
