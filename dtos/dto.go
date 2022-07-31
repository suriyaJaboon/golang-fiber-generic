package dtos

type Ok struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var OK = &Ok{
	Code:    "ok",
	Message: "successfully",
}
