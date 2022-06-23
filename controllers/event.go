package controller

import (
	"fmt"
	"misobo/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// employees data representation
type event struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Scheduled    string `json:"Scheduled"`
	Attend       []int  `json:"attend"`
	Accomodation []int  `json:"accomodation"`
}

// GetEvents responds with the list of all events as JSON.
func GetEvents(c *gin.Context) {
	var events = []event{}

	rows, err := utils.DB.Query(`select id, name, scheduled, attend, accomodation from "Events"`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var eventData event

		err = rows.Scan(&eventData.ID, &eventData.Name, &eventData.Scheduled, &eventData.Attend, &eventData.Accomodation)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		events = append(events, eventData)
		fmt.Println(eventData.Name, eventData.Scheduled)
	}
	c.IndentedJSON(http.StatusOK, events)
}

// GetEvent responds with a single event as JSON.
func GetEvent(c *gin.Context) {
	eventId := c.Param("event_id")

	rows, err := utils.DB.Query(`select id, name, scheduled, attend, accomodation from "Events" where id = $1`, eventId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var eventData event

	err = rows.Scan(&eventData.ID, &eventData.Name, &eventData.Scheduled, &eventData.Attend, &eventData.Accomodation)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(eventData.Name, eventData.Scheduled)
	c.IndentedJSON(http.StatusOK, eventData)
}

// GetEmployeesAttendingEvent responds with the list of all events as JSON.
func GetEmployeesAttendingEvent(c *gin.Context) {
	query := c.Request.URL.Query()
	eventId := c.Param("event_id")

	var arrayFormulation string

	employeeIds, err := GetEmployeeIds()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	for i := 0; i < len(employeeIds); i++ {
		if i+1 < len(arrayFormulation) {
			arrayFormulation += strconv.FormatInt(int64(employeeIds[i]), 10) + ","
		} else {
			arrayFormulation += strconv.FormatInt(int64(employeeIds[i]), 10)
		}
	}

	rows, err := utils.DB.Query(`select * from "Events" where attend && '{$1}' and id = $2`, arrayFormulation, eventId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var eventData event

	err = rows.Scan(&eventData.ID, &eventData.Name, &eventData.Scheduled, &eventData.Attend, &eventData.Accomodation)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	if query["accomodation"][0] == "1" {
		employeeIds = utils.RequireAccomodation(eventData.Attend, eventData.Accomodation)
	} else {
		employeeIds = utils.DontRequireAccomodation(eventData.Accomodation, eventData.Attend)
	}
	fmt.Println(eventData.Name, eventData.Scheduled)

	c.IndentedJSON(http.StatusOK, employeeIds)
}
