package model

type Collection struct {
	Name   string  `json:"name"`
	Routes []Route `json:"routes"`
	Groups []Group `json:"groups"`
}

type Group struct {
	Name   string  `json:"name"`
	Routes []Route `json:"routes"`
}

type Route struct {
	Method     string            `json:"method"`
	Path       string            `json:"path"`
	Headers    map[string]string `json:"headers,omitempty"`
	StatusCode int               `json:"statusCode,omitempty"`
	Body       any               `json:"body,omitempty"`
}

func (r *Route) GetStatusCode() int {
	if r.StatusCode == 0 {
		return 200
	}
	return r.StatusCode
}
