package model

type CollectionsRequest struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type CollectionRequest struct {
	Method      string `json:"method"`
	Body        string `json:"body"`
	ContentType string `json:"contentType"`
}
