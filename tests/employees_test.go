package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	controller "misobo/controllers"
	"misobo/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetEmloyeesHandler(t *testing.T) {
	utils.DatabaseConnection()

	path := "/employees"

	router := gin.Default()
	router.GET(path, controller.GetEmployees)
	req := httptest.NewRequest(http.MethodGet, path, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddEmloyeesHandlerCorrectFormat(t *testing.T) {
	utils.DatabaseConnection()

	path := "/employees"

	postBody := map[string]interface{}{
		"id":       "6",
		"name":     "Gustav",
		"gender":   "male",
		"birthday": "1992-12-11",
	}

	body, _ := json.Marshal(postBody)
	router := gin.Default()
	router.POST(path, controller.AddEmployees)
	req := httptest.NewRequest(http.MethodPost, path, bytes.NewReader(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	assert.Equal(t, http.StatusCreated, w.Code)

	//restoring the database state
	path = "/employees/6"
	router = gin.Default()
	router.DELETE("/employees/:employee_id", controller.DeleteEmployee)
	req = httptest.NewRequest(http.MethodDelete, path, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())
	fmt.Println("reponse", w.Body.String())
}

func TestAddEmloyeesHandlerWrongFormat(t *testing.T) {
	utils.DatabaseConnection()

	path := "/employees"

	postBody := map[string]interface{}{
		"id":     "6",
		"name":   "Gustav",
		"gender": "male",
	}

	body, _ := json.Marshal(postBody)
	router := gin.Default()
	router.POST(path, controller.AddEmployees)
	req := httptest.NewRequest(http.MethodPost, path, bytes.NewReader(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddEmloyeesHandlerWithEmptyBodyField(t *testing.T) {
	utils.DatabaseConnection()

	path := "/employees"

	postBody := map[string]interface{}{
		"id":       "6",
		"name":     "Gustav",
		"gender":   "",
		"birthday": "1992-12-11",
	}

	body, _ := json.Marshal(postBody)
	router := gin.Default()
	router.POST(path, controller.AddEmployees)
	req := httptest.NewRequest(http.MethodPost, path, bytes.NewReader(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateEmloyeesHandler(t *testing.T) {
	utils.DatabaseConnection()

	path := "/employees/:employee_id"

	postBody := map[string]interface{}{
		"id":       "2",
		"name":     "Devin",
		"gender":   "male",
		"birthday": "1992-12-10",
	}

	body, _ := json.Marshal(postBody)
	router := gin.Default()
	router.PUT(path, controller.UpdateEmployee)
	req := httptest.NewRequest(http.MethodPut, "/employees/2", bytes.NewReader(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdateEmloyeesHandlerWithWrongEmployeeId(t *testing.T) {
	utils.DatabaseConnection()

	path := "/employees/:employee_id"

	postBody := map[string]interface{}{
		"id":       "2",
		"name":     "Devin",
		"gender":   "male",
		"birthday": "1992-12-10",
	}

	body, _ := json.Marshal(postBody)
	router := gin.Default()
	router.PUT(path, controller.UpdateEmployee)
	req := httptest.NewRequest(http.MethodPut, "/employees/4", bytes.NewReader(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
