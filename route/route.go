package route

import "github.com/gin-gonic/gin"
import controller "golang-gin-mongo/controllers"

func MainRoute() {
	router := gin.Default()
	stringApi := "api"
	v1 := router.Group(stringApi + "/v1")
	{
		v1.GET("/student", controller.GetAllStudent)
		v1.DELETE("/student/delete/:first_name", controller.DeleteStudent)
		v1.POST("/student", controller.StoreStudent)
	}
	router.Run()
}
