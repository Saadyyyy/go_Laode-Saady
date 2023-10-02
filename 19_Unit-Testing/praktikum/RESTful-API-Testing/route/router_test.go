package route

import (
	"Laode_Saady/19_Unit-Testing/praktikum/RESTful-API-Testing/controller"
	"Laode_Saady/19_Unit-Testing/praktikum/RESTful-API-Testing/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestGetAllUser(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Define a recorder f
	rec := httptest.NewRecorder()

	// Define a request with the HTTP method "GET" and the path "/users"
	req := httptest.NewRequest(http.MethodGet, "/users", nil)

	// Create an Echo context from the request and recorder
	eContext := e.NewContext(req, rec)

	// Mock the JwtMiddleware
	e.Use(middleware.JwtMiddleware())

	// Call the controller function (GetUsersController) using the Echo context
	err := controller.GetUsersController(eContext)
	if err != nil {
		t.Errorf("Error in GetUsersController: %v", err)
	}

	// Perform assertions on the response, status code, or other aspects if needed
	// For example, you can check the status code:
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rec.Code)
	}
}
