package psql

import "misobo/utils"

// employees data representation
type employee struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}

func Find() ([]employee, error) {
	var employees = []employee{}

	rows, err := utils.DB.Query(`select id, name, gender, birthday from "Employees"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var newEmployee employee
		err = rows.Scan(
			&newEmployee.ID,
			&newEmployee.Name,
			&newEmployee.Gender,
			&newEmployee.Birthday,
		)

		if err != nil {
			return nil, err
		}
		employees = append(employees, newEmployee)
	}

	return employees, nil
}
