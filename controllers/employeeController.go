package controllers

import (
	appErrors "GoLayeredCRUD/errors"
	"GoLayeredCRUD/models"
	"GoLayeredCRUD/services"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Timestamp  string      `json:"timestamp"`
	Payload    interface{} `json:"data"`
}

func GetEmployees(ctx *gin.Context) {

	employees, err := services.GetEmployee(ctx)

	if err != nil && err.Error() == "no employee found" {
		ctx.JSON(http.StatusOK, Response{
			Status:     "error",
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Timestamp:  time.Now().Format(time.RFC3339),
			Payload:    employees,
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Status:     "succes",
		StatusCode: 200,
		Message:    "Successfully fetched all employees",
		Timestamp:  time.Now().Format(time.RFC3339),
		Payload:    employees,
	})

}

func GetEmployeByID(ctx *gin.Context) {
	employeeId := ctx.Param("id")
	results, err := services.GetEmployeByID(employeeId)

	if errors.Is(err, appErrors.ErrEmployeeNotFound) {
		ctx.JSON(http.StatusNotFound, Response{
			Status:     "error",
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Timestamp:  time.Now().Format(time.RFC3339),
			Payload:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Status:     "success",
		StatusCode: 200,
		Message:    "Successfully fetched employees",
		Timestamp:  time.Now().Format(time.RFC3339),
		Payload:    results,
	})

}

func StoreEmployee(ctx *gin.Context) {
	var reqPayload models.Employee

	// validation JSON request
	if err := ctx.ShouldBind(&reqPayload); err != nil {
		fmt.Println(&reqPayload)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":     "error",
			"statusCode": 400,
		})
		return
	}

	_, err := services.StoreEmployee(reqPayload.ID, reqPayload)
	if errors.Is(err, appErrors.ErrEmployeeAlreadyExists) {
		ctx.JSON(http.StatusConflict, Response{
			Status:     "error",
			StatusCode: http.StatusConflict,
			Message:    appErrors.ErrEmployeeAlreadyExists.Error(),
			Timestamp:  time.Now().Format(time.RFC3339),
			Payload:    nil,
		})
		return
	}

	if errors.Is(err, appErrors.ErrEmployeeIsRequired) {
		ctx.JSON(http.StatusBadRequest, Response{
			Status:     "error",
			StatusCode: http.StatusBadRequest,
			Message:    appErrors.ErrEmployeeIsRequired.Error(),
			Timestamp:  time.Now().Format(time.RFC3339),
			Payload:    nil,
		})
		return
	}
	if errors.Is(err, appErrors.ErrInvalidIdEmployee) {
		ctx.JSON(http.StatusBadRequest, Response{
			Status:     "error",
			StatusCode: http.StatusBadRequest,
			Message:    appErrors.ErrInvalidIdEmployee.Error(),
			Timestamp:  time.Now().Format(time.RFC3339),
			Payload:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Status:     "success",
		StatusCode: http.StatusCreated,
		Message:    "Succesfully created employe",
		Timestamp:  time.Now().Format(time.RFC3339),
		Payload:    reqPayload,
	})

}

func UpdateEmployee(ctx *gin.Context) {
	idEmployee := ctx.Param("id")

	// Bind JSON body ke struct Employee
	var updatedEmployee models.Employee
	if err := ctx.ShouldBind(&updatedEmployee); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Status:     "error",
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Data",
			Timestamp:  time.Now().Format(time.RFC3339),
			Payload:    nil,
		})
		return
	}

	if updatedEmployee.ID == "" {
		ctx.JSON(http.StatusBadRequest, Response{
			Status:     "error",
			StatusCode: http.StatusBadRequest,
			Message:    "Employee id is required",
			Timestamp:  time.Now().Format(time.RFC3339),
			Payload:    nil,
		})
		return
	}

	results := services.UpdateEmployee(idEmployee, updatedEmployee)
	if results {
		ctx.JSON(http.StatusOK, Response{
			Status:     "success",
			StatusCode: http.StatusOK,
			Message:    "Succesfully Updated employee",
			Timestamp:  time.Now().Format(time.RFC3339),
			Payload:    updatedEmployee,
		})
	} else {
		ctx.JSON(http.StatusNotFound, Response{
			Status:     "error",
			StatusCode: http.StatusNotFound,
			Message:    "Employee Not Found",
			Timestamp:  time.Now().Format(time.RFC3339),
			Payload:    nil,
		})
	}

}

func DeleteEmployee(c *gin.Context) {
	idEmployee := c.Param("id")
	results, err := services.DeleteEmployee(idEmployee)

	if errors.Is(err, appErrors.ErrEmployeeNotFound) {
		c.JSON(http.StatusNotFound, Response{
			Status:     "error",
			StatusCode: 404,
			Message:    err.Error(),
			Timestamp:  time.Now().Format(time.RFC3339),
			Payload:    nil,
		})
		return
	}

	message := fmt.Sprintf("Employee successfully deleted , with ID :%s", results)
	c.JSON(http.StatusOK, Response{
		Status:     "success",
		StatusCode: 200,
		Message:    message,
		Timestamp:  time.Now().Format(time.RFC3339),
		Payload:    nil,
	})
}
