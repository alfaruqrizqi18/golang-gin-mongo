package controllers

import (
	config "golang-gin-mongo/config"
	structs "golang-gin-mongo/structs"
	"log"
	"net/http"
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
	first_name := c.DefaultQuery("first_name", "")

	if strings.EqualFold(first_name, "") {
		var student []structs.Student

		error := studentCollection.Find(nil).Sort("-$natural").All(&student)
		if error != nil {
			log.Fatal(error.Error())
		}

		if len(student) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "student not found",
				"data":    nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "student data",
				"data":    student,
			})
		}

	} else {

		var student []structs.Student
		// first_name := c.DefaultQuery("first_name", "")
		error := studentCollection.
			Find(bson.M{"first_name": bson.RegEx{Pattern: first_name, Options: "i"}}). // options i adalah untuk ignore case
			All(&student)
		if error != nil {
			log.Fatal(error.Error())
		}
		if len(student) > 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "similar student",
				"data":    student,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "similar student not found",
				"data":    nil,
			})
		}
	}

}

func GetSingleStudentById(c *gin.Context) {
	var detailStudent = structs.Student{}
	if bson.IsObjectIdHex(c.Param("id")) { // validate id is ObjectIdHex
		id := bson.ObjectIdHex(c.Param("id"))
		error := studentCollection.FindId(id).One(&detailStudent)

		if error != nil {
			log.Fatal(error.Error())
		}

		if detailStudent != (structs.Student{}) {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "detail student found",
				"data":    detailStudent,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "student not found",
				"data":    nil,
			})
		}

	} else { // validate false
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "id is not valid",
			"data":    nil,
		})
	}

}

func StoreStudent(c *gin.Context) {
	id := bson.NewObjectId()
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	created_at := c.PostForm("created_at")

	storeStudent := structs.Student{id, first_name, last_name, created_at}
	error := studentCollection.Insert(&storeStudent)

	if error != nil {
		log.Fatal(error.Error())
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "new student created successfully",
		"data":    storeStudent,
	})
}

func DeleteStudent(c *gin.Context) {
	if bson.IsObjectIdHex(c.Param("id")) {
		var beforeDelete, afterDelete []structs.Student

		id := bson.ObjectIdHex(c.Param("id"))
		studentCollection.Find(nil).Sort("-$natural").All(&beforeDelete)
		error := studentCollection.RemoveId(id)
		studentCollection.Find(nil).Sort("-$natural").All(&afterDelete)
		if error != nil {
			log.Fatal(error)
			c.JSON(http.StatusNotFound, gin.H{
				"status":       http.StatusNotFound,
				"message":      "student not found",
				"beforeDelete": nil,
				"afterDelete":  nil,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"status":       http.StatusOK,
			"message":      "1 student removed successfully",
			"beforeDelete": beforeDelete,
			"afterDelete":  afterDelete,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "id is not valid",
		})
	}
}
