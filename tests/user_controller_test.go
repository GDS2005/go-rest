package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example/controllers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Create a new Gin engine instance
	r := gin.Default()

	// Define routes
	r.GET("/users", controllers.GetUsers)

	// Serve HTTP
	r.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body (assuming it's JSON)
	// assert.Contains(t, rr.Body.String(), `"users":`)
}
