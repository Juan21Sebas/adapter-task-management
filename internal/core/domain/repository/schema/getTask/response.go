package gettask

type TaskResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
