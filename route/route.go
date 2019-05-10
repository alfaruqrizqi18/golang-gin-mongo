package route

import "github.com/gin-gonic/gin"
import controller "golang-gin-mongo/controllers"

func MainRoute() {
	router := gin.Default()
	stringApi := "api"
	v1 := router.Group(stringApi + "/v1")
	{
		v1.POST("/student", controller.StoreStudent)
	}

	router.Run()
}
