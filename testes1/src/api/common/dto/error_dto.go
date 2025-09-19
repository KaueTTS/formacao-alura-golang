package api_common

type ErrorDto struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Code    int    `json:"code"`
}
