package response

type Response struct {
	Data  interface{} `json:"data"`
	Meta  Meta        `json:"meta"`
	Error interface{} `json:"error,omitempty"`
}
