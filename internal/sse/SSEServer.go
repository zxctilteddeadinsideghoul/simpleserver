package sse

import (
	"sync"
)

type Server struct {
	clients map[chan string]bool
	mu      sync.Mutex
}

func NewSSEServer() *Server {
	return &Server{
		clients: make(map[chan string]bool),
	}
}

func (s *Server) AddClient(client chan string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clients[client] = true
}

func (s *Server) RemoveClient(client chan string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.clients, client)
	close(client)
}

func (s *Server) Broadcast(message string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for client := range s.clients {
		client <- message
	}
}

var Instance *Server

func Init() {
	Instance = NewSSEServer()
}
