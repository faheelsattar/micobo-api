package tests

import (
	controller "misobo/controllers"
	"misobo/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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
}
