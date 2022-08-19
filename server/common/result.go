package common

type Result struct {
	Code    int         `json:"code"`
	MESSAGE string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Error(code int, msg string) *Result {
	return &Result{
		Code:    code,
		MESSAGE: msg,
	}
}

func Success(code int, data interface{}) *Result {
	return &Result{
		Code: code,
		Data: data,
	}
}
