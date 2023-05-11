package tests

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHandler_ConfigurePrey(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		// Arrange
		router := CreateServer()
		request, response := CreateRequestTest("PUT", "/v1/prey", `{"speed": 10.0}`)
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedResponse := `{"success":true}`

		// Act
		router.ServeHTTP(response, request)

		// Assert
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expectedHeader, response.Header())
		assert.Equal(t, expectedResponse, response.Body.String())
	})

	t.Run("Invalid Field Type", func(t *testing.T) {
		// Arrange
		router := CreateServer()
		request, response := CreateRequestTest("PUT", "/v1/prey", `{"speed": "hello_world"}`)
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedResponse := `{"success":false}`

		// Act
		router.ServeHTTP(response, request)

		// Assert
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expectedHeader, response.Header())
		assert.Equal(t, expectedResponse, response.Body.String())
	})
}

func TestHandler_ConfigureShark(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		// Arrange
		router := CreateServer()
		request, response := CreateRequestTest("PUT", "/v1/shark", `{"x_position": 10.0, "y_position": 10.0, "speed": 10.0}`)
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedResponse := `{"success":true}`

		// Act
		router.ServeHTTP(response, request)

		// Assert
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expectedHeader, response.Header())
		assert.Equal(t, expectedResponse, response.Body.String())
	})

	t.Run("Invalid Field Type", func(t *testing.T) {
		// Arrange
		router := CreateServer()
		request, response := CreateRequestTest("PUT", "/v1/shark", `{"x_position": "hello_world", "y_position": 10.0, "speed": 10.0}`)
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedResponse := `{"success":false}`

		// Act
		router.ServeHTTP(response, request)

		// Assert
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expectedHeader, response.Header())
		assert.Equal(t, expectedResponse, response.Body.String())
	})
}

func TestHandler_SimulateHunt(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		// Arrange
		router := CreateServer()
		request, response := CreateRequestTest("POST", "/v1/simulate", "")
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedResponse := `{"success":true,"message":"Simulation completed","time":3}`

		// Act
		router.ServeHTTP(response, request)

		// Assert
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expectedHeader, response.Header())
		assert.Equal(t, expectedResponse, response.Body.String())
	})
}
