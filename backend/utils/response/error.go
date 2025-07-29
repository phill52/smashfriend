package response

type Error struct {
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"code"`
}
