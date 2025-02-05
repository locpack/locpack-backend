package dtos

type Meta struct {
	Success bool `json:"success"`
}

type Error struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type ResponseWrapper[T any] struct {
	Data   T       `json:"data"`
	Meta   Meta    `json:"meta"`
	Errors []Error `json:"errors"`
}
