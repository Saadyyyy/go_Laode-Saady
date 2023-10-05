package controller

// import (
// 	"testing"

// 	"net/http"
// 	"net/http/httptest"

// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetAllUsers(t *testing.T) {
// 	// Create a new Echo instance for testing
// 	e := echo.New()

// 	// Create a mock userCase with a GetAllUsers method that returns a sample userList and no error
// 	// mockUserCase := &MockUserCase{} // You should create this mock implementation

// 	// Create a request to simulate a GET request to /users (or your desired endpoint)
// 	req := httptest.NewRequest(http.MethodGet, "/users", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	// Create the UserController and set its userCase to the mock
// 	uc := UserController{
// 		userCase: mockUserCase,
// 	}

// 	// Invoke the GetAllUsers function
// 	if assert.NoError(t, uc.GetAllUsers(c)) {
// 		// Check the response status code (200 OK)
// 		assert.Equal(t, http.StatusOK, rec.Code)

// 		// Check the response body (you can parse JSON and validate its content)
// 		// For example, if you expect a JSON response like {"message": "Success getting all users", "users": [userList]}
// 		expected := `{"message":"Success getting all users","users":[/* userList JSON here */]}`
// 		assert.JSONEq(t, expected, rec.Body.String())
// 	}
// }
