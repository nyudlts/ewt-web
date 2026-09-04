package api_v0

type IndexResponse struct {
	Message     string   `json:"message"`
	EwtHome     string   `json:"ewt_home"`
	Directories []string `json:"directories"`
}
