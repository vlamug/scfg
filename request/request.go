package request

// GetRequest is a struct of request to load config
type GetRequest struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
