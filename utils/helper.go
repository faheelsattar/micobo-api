package utils

func RequireAccomodation(attenders []int, accomodation []int) []int {
	var employees []int
	hash := make(map[int]bool)
	for _, e := range attenders {
		hash[e] = true
	}
	for _, e := range accomodation {
		// If elements present in the hashmap then append intersection list.
		if hash[e] {
			employees = append(employees, e)
		}
	}
	return employees
}

func DontRequireAccomodation(attenders []int, accomodation []int) []int {
	var employees []int
	hash := make(map[int]bool)
	for _, e := range accomodation {
		hash[e] = true
	}
	for _, e := range attenders {
		// If elements absent in the hashmap then append intersection list.
		if !hash[e] {
			employees = append(employees, e)
		}
	}
	return employees
}
