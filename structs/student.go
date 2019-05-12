package structs

import "gopkg.in/mgo.v2/bson"

type Student struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id" form:"id"`
	FirstName string        `bson:"first_name" json:"first_name" form:"first_name"`
	LastName  string        `bson:"last_name" json:"last_name" form:"last_name"`
	CreatedAt string        `bson:"created_at" json:"created_at" form:"created_at"`
}
