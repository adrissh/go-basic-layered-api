package services

import (
	appErrors "GoLayeredCRUD/errors"
	"GoLayeredCRUD/models"
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
)

var employees = []models.Employee{
	{ID: "EMP001", Name: "John Doe", Age: 30, Position: "Software Engineer", Salary: 75000.50, HireDate: "2022-04-15"},
	{ID: "EMP002", Name: "Jane Smith", Age: 28, Position: "Data Analyst", Salary: 68000.00, HireDate: "2021-06-23"},
	{ID: "EMP003", Name: "Michael Brown", Age: 35, Position: "Product Manager", Salary: 85000.75, HireDate: "2019-03-12"},
	{ID: "EMP004", Name: "Sarah Johnson", Age: 40, Position: "HR Manager", Salary: 72000.00, HireDate: "2017-11-05"},
	{ID: "EMP005", Name: "David Williams", Age: 25, Position: "Junior Developer", Salary: 55000.00, HireDate: "2023-01-10"},
}

// var employees = []models.Employee{}

func GetEmployee(ctx *gin.Context) ([]models.Employee, error) {
	if len(employees) < 1 {
		return []models.Employee{}, fmt.Errorf("no employee found")
	}
	return employees, nil
}

func GetEmployeByID(employeeId string) (models.Employee, error) {
	for _, emp := range employees {
		if employeeId == emp.ID {
			return emp, nil
		}
	}
	return models.Employee{}, fmt.Errorf("%w : employeID :%s", appErrors.ErrEmployeeNotFound, employeeId)

}

func StoreEmployee(idEmployee string, employee models.Employee) (models.Employee, error) {

	// check idEmploye, if same with slice employee then reject
	for _, employee := range employees {
		if idEmployee == employee.ID {
			return models.Employee{}, fmt.Errorf("%w", appErrors.ErrEmployeeAlreadyExists)
		}
	}

	// validation for not empty employee id
	if idEmployee == "" {
		return models.Employee{}, fmt.Errorf("%w", appErrors.ErrEmployeeIsRequired)
	}
	// validation format employe id
	match, _ := regexp.MatchString("EMP+[0-9]", idEmployee)
	if !match {
		return models.Employee{}, fmt.Errorf("%w", appErrors.ErrInvalidIdEmployee)
	}

	employees = append(employees, employee)
	return models.Employee{}, nil

}

func UpdateEmployee(idEmployee string, updatedEmployee models.Employee) bool {
	for idx, employee := range employees {
		if employee.ID == idEmployee { // if data found , then update data
			employees[idx] = updatedEmployee
			return true
		}
	}
	return false
}

func DeleteEmployee(idEmployee string) (string, error) {

	for idx, emp := range employees {
		if emp.ID == idEmployee { // if employe found , then delete data
			index := idx
			employees = append(employees[:index], employees[index+1:]...)
			return idEmployee, nil
		}
	}

	return "", fmt.Errorf("%w: '%s'", appErrors.ErrEmployeeNotFound, idEmployee)

}
