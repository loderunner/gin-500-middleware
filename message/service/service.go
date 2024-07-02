package service

import (
	"errors"
	"slices"
)

type MessageService struct {
	users []string
}

var ErrUserNotFound = errors.New("user not found")

func NewMessageService() *MessageService {
	return &MessageService{
		users: []string{
			"alice",
			"bob",
		},
	}
}

func (s *MessageService) Send(message string, to string) error {
	if !slices.Contains(s.users, to) {
		return ErrUserNotFound
	}
	return nil
}
