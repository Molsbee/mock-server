package model

type Collection struct {
	Name   string  `json:"name"`
	Routes []Route `json:"routes"`
}

type Route struct {
	Method     string            `json:"method"`
	Path       string            `json:"path"`
	Headers    map[string]string `json:"headers"`
	StatusCode int               `json:"statusCode"`
	Body       any               `json:"body"`
}
