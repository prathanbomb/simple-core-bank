package routes

const APP_CTX_KEY = "appCtx"

type Response struct {
	Code       uint64      `json:"code"`
	Message    string      `json:"message"`
	Pagination interface{} `json:"pagination,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}
