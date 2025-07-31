package response

type Response struct {
	Data  interface{} `json:"data"`
	Meta  Meta        `json:"meta"`
	Error interface{} `json:"error,omitempty"`
}

type Meta struct {
	Status     int         `json:"status"`
	Message    string      `json:"message"`
	Pagination interface{} `json:"pagination,omitempty"`
}

type Error struct {
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"code"`
}
