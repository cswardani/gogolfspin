package model

import "gopkg.in/mgo.v2/bson"

// Field table
type Field struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	NAME       string        `bson:"NAME" json:"NAME"`
	LOCATION   string        `bson:"LOCATION" json:"LOCATION"`
	FACILITIES string        `bson:"FACILITIES" json:"FACILITIES"`
	PRICE      string        `bson:"PRICE" json:"PRICE"`
	WEBSITE    string        `bson:"WEBSITE" json:"WEBSITE"`
	PHONE      string        `bson:"PHONE" json:"PHONE"`
	ABOUT      string        `bson:"ABOUT" json:"ABOUT"`
	IMAGE      string        `bson:"IMAGE" json:"IMAGE"`
}
