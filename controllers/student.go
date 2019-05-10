package controllers

import (
	config "golang-gin-mongo/config"
	structs "golang-gin-mongo/structs"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
)

var studentCollection *mgo.Collection
var session *mgo.Session

func init() {
	session, _ = config.Connect()
	studentCollection = session.DB(config.DatabaseName).C("student")
}

func GetAllStudent(c *gin.Context) {

}

func StoreStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	created_at := c.PostForm("created_at")
	defer session.Close()

	storeStudent := structs.Student{id, first_name, last_name, created_at}
	error := studentCollection.Insert(&storeStudent)

	if error != nil {
		log.Fatal(error)
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "new student created successfully",
	})
}
