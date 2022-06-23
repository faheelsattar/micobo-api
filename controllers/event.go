package controller

import (
	"fmt"
	"misobo/utils"
	"net/http"

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
