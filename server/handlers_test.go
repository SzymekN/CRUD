package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestHelloWorld(t *testing.T) {
	// Build our expected body
	// body := map[string]string{
	// 	"message": "Hello, World!",
	// }

	router := SetupRouter()

	w := performRequest(router, "GET", "/")

	assert.Equal(t, http.StatusOK, w.Code)

	objectString := `{
        "id": 2001,
        "name": "U-name",
        "email": "gool@gool.com"
    }`
	var unknown map[string]interface{}
	err := json.Unmarshal([]byte(objectString), &unknown)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(unknown, unknown["name"])

	var response map[string]interface{}
	// err = json.Unmarshal([]byte(w.Body.String()), &response)
	// if err != nil {
	// fmt.Println(err)
	// }

	// Make some assertions on the correctness of the response.
	value, exists := response["message"]
	fmt.Println((w.Body.String()))
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "{\"message\":\"Hello World!\"}", value)

}
