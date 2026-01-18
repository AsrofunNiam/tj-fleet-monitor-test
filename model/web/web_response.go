package web

type WebResponse struct {
	Success   bool        `json:"success"`
	TotalData int         `json:"total_data"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}
