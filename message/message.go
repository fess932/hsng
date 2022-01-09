package message

import (
	"github.com/fess932/hsng/graph/model"
)

func New(repo Repo) *Messager {
	m := &Messager{
		messages:  []*model.Message{},
		reciver:   make(chan *model.Message),
		listeners: map[string]chan *model.Message{},
		repo:      repo,
	}

	go m.messageService() // run message listener

	return m
}

type Messager struct {
	reciver   chan *model.Message
	listeners map[string]chan *model.Message
	messages  []*model.Message
	repo      Repo
}

type Repo interface {
	ReadAllMessages() []*model.Message
	SaveMessage(message *model.Message)
}

func (m *Messager) messageService() {

	for {
		select {
		case msg := <-m.reciver:
			m.messages = append(m.messages, msg)

			// broadcast message to all subscibers
			for _, v := range m.listeners {
				v <- msg
			}
		}
	}
}

func (m Messager) Subscribe(userID string) <-chan *model.Message {
	ch := make(chan *model.Message)
	m.listeners[userID] = ch
	return ch
}

func (m Messager) Unsubscribe(userID string) {
	close(m.listeners[userID])
	delete(m.listeners, userID)
}

func (m Messager) GetMessages() []*model.Message {
	return m.repo.ReadAllMessages()
}

func (m Messager) SendMessage(message *model.Message) {
	m.reciver <- message
	m.repo.SaveMessage(message)
}
