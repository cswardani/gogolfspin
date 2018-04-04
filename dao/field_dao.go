package dao

import (
	"log"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"../model/"
)

// FieldsDAO type
type FieldsDAO struct {
	Server string
	Database string
}

var db *mgo.Database

// COLLECTION of fields
const ( COLLECTION = "fields")

// Connect establish
func (f *FieldsDAO) Connect() {
	session, err := mgo.Dial(f.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(f.Database)
}

// FindAll fields
func (f *FieldsDAO) FindAll() ([]Field, error){
	var field []Field
	err := db.C(COLLECTION).Find(bson.M{}).All(&field)
}

//FindByID fields
func (f * FieldsDAO) FindByID(id int) (Field, error){
	var field Field
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&field)
	return field, err
}

// Insert a field into database
func (f *FieldsDAO) Insert(field Field) error {
	err := db.C(COLLECTION).Insert(&field)
	return err
}

// Delete an existing field
func (f *FieldsDAO) Delete(field Field) error {
	err := db.C(COLLECTION).Remove(&field)
	return err
}

// Update an existing field
func (f *FieldsDAO) Update(field Field) error {
	err := db.C(COLLECTION).UpdateId(field.ID, &field)
	return err
}