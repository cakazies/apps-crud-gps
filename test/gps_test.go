package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/local/app-gps/application/api"
	"github.com/local/app-gps/routes"
)

type testCase struct {
	name         string
	input        string
	expectedData string
	expectedCode int
	path         string
	handler      routes.HandlerFunc
	query        string
}

type Response struct {
	Response Rest                   `json:"response"`
	Data     map[string]interface{} `json:"data,omitempty"`
}

type Rest struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}

func TestGetRoom(t *testing.T) {
	tasks := []testCase{
		{
			name:         "Testing with id",
			input:        Cfg.ID,
			expectedData: Cfg.ID,
			expectedCode: http.StatusOK,
			path:         "api/gps",
			handler:      api.GetGPS,
			query:        "",
		},
		{
			name:         "Testing with random id",
			input:        "9897",
			expectedData: "<nil>", // because not value
			expectedCode: http.StatusBadRequest,
			path:         "api/gps",
			handler:      api.GetGPS,
			query:        "",
		},
		{
			name:         "Testing with variable",
			input:        "abc",
			expectedData: "<nil>", // because not value
			expectedCode: http.StatusBadRequest,
			path:         "api/gps",
			handler:      api.GetGPS,
			query:        "",
		},
	}

	for _, tc := range tasks {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("http://%s/%s/%s", Cfg.URL, tc.path, tc.input)
			resp := getRequest(url, "", tc.handler)
			assert.Equal(t, resp.Code, tc.expectedCode, "Expedted Code is Wrong")
			buf := resp.Body.Bytes()
			var respData Response
			if err := json.Unmarshal(buf, &respData); err != nil {
				t.Error("Can not parsing response testing. Error :", err)
			}
			getData := fmt.Sprintf("%v", respData.Data["id"])
			// log.Println(getData)
			assert.Equal(t, getData, tc.expectedData, "Expedted Data is Wrong")
		})
	}
}
