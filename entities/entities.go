package entities

// Repository represent the database functions
type Repository interface {
	Close()

	FindEmployee() ([]*Employee, error)
	FindEmployeeIds() ([]int, error)
	EmployeeExists(employeeId string) bool
	CreateEmployee(employee *Employee) error
	UpdateEmployee(employee *Employee, employeeId string) error
	DeleteEmployee(employeeId string) error

	FindEvents() ([]*Event, error)
	FindSingleEvent(eventId string) (*Event, error)
	FindEmployeesAttendingEvent(employeeIdsString string, eventId string) (*Event, error)
}

// Employee data representation
type Employee struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}

// Event data representation
type Event struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Scheduled    string `json:"Scheduled"`
	Attend       []int  `json:"attend"`
	Accomodation []int  `json:"accomodation"`
}
