package controllers

import (
	config "golang-gin-mongo/config"
	structs "golang-gin-mongo/structs"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var studentCollection *mgo.Collection
var session *mgo.Session

func init() {
	session, _ = config.Connect()
	studentCollection = session.DB(config.DatabaseName).C("student")
}

func GetAllStudent(c *gin.Context) {
	var student []structs.Student
	first_name := c.DefaultQuery("first_name", "")

	if strings.EqualFold(first_name, "") {

		error := studentCollection.Find(nil).All(&student)
		if error != nil {
			log.Fatal(error)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "student data",
			"data":    student,
		})

	} else {

		var detailStudent = structs.Student{}
		first_name := c.DefaultQuery("first_name", "")
		error := studentCollection.Find(bson.M{"first_name": &first_name}).One(&detailStudent)
		if error != nil {
			log.Fatal(error)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "detail student",
			"data":    detailStudent,
		})
	}

}

func StoreStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	created_at := c.PostForm("created_at")

	storeStudent := structs.Student{id, first_name, last_name, created_at}
	error := studentCollection.Insert(&storeStudent)

	if error != nil {
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "new student created successfully",
	})
}

func DeleteStudent(c *gin.Context) {
	var beforeDelete, afterDelete []structs.Student

	first_name := c.Param("first_name")
	studentCollection.Find(nil).All(&beforeDelete)
	error := studentCollection.Remove(bson.M{"first_name": first_name})
	studentCollection.Find(nil).All(&afterDelete)
	if error != nil {
		log.Fatal(error)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"message":      "1 student removed successfully",
		"beforeDelete": beforeDelete,
		"afterDelete":  afterDelete,
	})
}
