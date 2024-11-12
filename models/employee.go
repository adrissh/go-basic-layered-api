package models

type Employee struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
	HireDate string  `json:"hireDate"`
}
