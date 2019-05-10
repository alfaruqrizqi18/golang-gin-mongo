package structs

type Student struct {
	Id        int    `bson:"id" json:"id" form:"id"`
	FirstName string `bson:"first_name" json:"first_name" form:"first_name"`
	LastName  string `bson:"last_name" json:"last_name" form:"last_name"`
	CreatedAt string `bson:"created_at" json:"created_at" form:"created_at"`
}
