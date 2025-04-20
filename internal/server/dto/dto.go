package dto

type Meta struct {
	Success bool `json:"success"`
}

type Error struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type ResponseWrapper struct {
	Data   any     `json:"data"`
	Meta   Meta    `json:"meta"`
	Errors []Error `json:"errors"`
}
