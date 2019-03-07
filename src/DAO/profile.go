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
		return nil, fmt.Errorf("Error to get all Profiles: %v", err)
	}
	return profiles, err
}

// GetByID get only Profile document
func (p *Profile) GetByID(id string) (Profile, error) {
	var profile Profile
	// You can use:
	// err := Db.C(collection).FindId(bson.ObjectIdHex(id)).One(&profile)
	err := Db.C(collection).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&profile)
	if err != nil {
		return profile, fmt.Errorf("Error to get only Profile: %v", err)
	}
	return profile, err
}

// Update the p (Profile) to other profile (Profile)
func (p *Profile) Update(profile *Profile) error {
	err := Db.C(collection).Update(p, profile)
	if err != nil {
		return fmt.Errorf("Error to update Profile: %v", err)
	}
	return err
}

// UpdateByID update the Profile with its id by p (Profile)
func (p *Profile) UpdateByID(id string) error {
	err := Db.C(collection).UpdateId(bson.ObjectIdHex(id), p)
	if err != nil {
		return fmt.Errorf("Error to update by id Profile: %v", err)
	}
	return err
}

// DeleteByID delete a Profile by ID
func (p *Profile) DeleteByID(id string) error {
	err := Db.C(collection).RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		return fmt.Errorf("Error to delete by id Profile: %v", err)
	}
	return err
}
