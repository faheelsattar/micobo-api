package entities

// Repository represent the database functions
type Repository interface {
	Close()
	FindEmployee() ([]*Employee, error)
	FindEmployeeIds() ([]int, error)
	EmployeeExists(employeeId string) (*Employee, error)
	CreateEmployee(employee *Employee) error
	UpdateEmployee(employee *Employee, employeeId string) error
	DeleteEmployee(employeeId string) error
}

// Employee data representation
type Employee struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}
