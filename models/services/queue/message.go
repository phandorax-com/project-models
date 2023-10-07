package service

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Message[T any] struct {
	Uuid       string            `json:"uuid"`
	DataTime   time.Time         `json:"date_time"`
	Aplication AplicationContext `json:"producer"`
	Message    T                 `json:"message"`
}

type AplicationContext struct {
	Name     string `json:"app_name"`
	Tracking string `json:"tracking"`
	UserID   string `json:"user_id"`
}

func SendMessage(aplication AplicationContext, value any) Message[any] {
	return Message[any]{
		Uuid:       uuid.NewString(),
		DataTime:   time.Now().UTC(),
		Aplication: aplication,
		Message:    value,
	}
}

func GetMessage(message string) Message[any] {

	data := Message[any]{}

	err := json.Unmarshal([]byte(message), &data)

	if err != nil {
		return data
	}

	return data
}
