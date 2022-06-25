package controller

import (
	"fmt"
	"misobo/psql"
	"misobo/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetEvents responds with the list of all events as JSON.
func GetEvents(c *gin.Context) {
	repo := &psql.Repository{Db: utils.DB}

	events, err := repo.FindEvents()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, events)
}

// GetEvent responds with a single event as JSON.
func GetEvent(c *gin.Context) {
	repo := &psql.Repository{Db: utils.DB}

	eventId := c.Param("event_id")

	event, err := repo.FindSingleEvent(eventId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, event)
}

// GetEmployeesAttendingEvent responds with the list of all events as JSON.
func GetEmployeesAttendingEvent(c *gin.Context) {
	repo := &psql.Repository{Db: utils.DB}

	query := c.Request.URL.Query()
	eventId := c.Param("event_id")

	var arrayFormulation string

	employeeIds, err := repo.FindEmployeeIds()

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

	event, err := repo.FindEmployeesAttendingEvent(arrayFormulation, eventId)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	if query["accomodation"][0] == "1" {
		employeeIds = utils.RequireAccomodation(event.Attend, event.Accomodation)
	} else {
		employeeIds = utils.DontRequireAccomodation(event.Accomodation, event.Attend)
	}
	fmt.Println(event.Name, event.Scheduled)

	c.IndentedJSON(http.StatusOK, employeeIds)
}
