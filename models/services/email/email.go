package service

type Message[T any] struct {
	To       string `json:"to"`
	Subject  string `json:"subject"`
	Template string `json:"template"`
	File     string `json:"file"`
	Data     T      `json:"data"`
}
