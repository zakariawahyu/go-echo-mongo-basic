package response

type WebResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Results interface{} `json:"results"`
}
