package model

type Collection struct {
	Name   string  `json:"name"`
	Routes []Route `json:"routes"`
}

type Route struct {
	Path       string `json:"path"`
	Method     string `json:"method"`
	StatusCode int    `json:"statusCode"`
	Body       any    `json:"body"`
}
