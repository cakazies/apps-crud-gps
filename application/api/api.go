package api

// Response REST API
type Response struct {
	Response Rest        `json:"response"`
	Data     interface{} `json:"data,omitempty"`
}

type Rest struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}
