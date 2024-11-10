package routes

import (
	"GoLayeredCRUD/controllers"

	"github.com/gin-gonic/gin"
)

func EmployeeRoutes(router *gin.Engine) {
	//  Simple group: v1
	v1 := router.Group("/api/v1")
	{
		v1.GET("/employees", controllers.GetEmployees)
		v1.GET("/employees/:id", controllers.GetEmployeByID)
		v1.POST("/employees", controllers.StoreEmployee)
		v1.PUT("/employees/:id", controllers.UpdateEmployee)
		v1.DELETE("/employees/:id", controllers.DeleteEmployee)
	}

}
