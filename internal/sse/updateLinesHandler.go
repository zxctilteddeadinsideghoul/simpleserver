package sse

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) UpdateLinesEventHandler(c *gin.Context) {
	// set sse headers
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	clientChan := make(chan string)
	s.AddClient(clientChan)
	defer s.RemoveClient(clientChan)

	// sent events
	for {
		select {
		case _, ok := <-clientChan:
			if !ok {
				return // chan closed
			}
			c.Writer.Flush()
		case <-c.Request.Context().Done():
			return // client close connection
		}
	}
}
