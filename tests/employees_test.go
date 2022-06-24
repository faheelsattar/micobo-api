package tests

import (
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

	expectedData := `{"data":[{"id":5,"name":"faheel sattar","gender":"male","birthday":"1998-01-01T00:00:00Z"},{"id":2,"name":"kevin","gender":"male","birthday":"1998-01-01T00:00:00Z"},{"id":3,"name":"Levin","gender":"female","birthday":"1998-01-01T00:00:00Z"}]}`
	path := "/employees"

	router := gin.Default()
	router.GET(path, controller.GetEmployees)
	req := httptest.NewRequest(http.MethodGet, path, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("status: %d", w.Code)
	t.Logf("response: %s", w.Body.String())

	fmt.Println("reponse", expectedData == w.Body.String())
	assert.Equal(t, w.Body.String(), expectedData)

}
