package logger

import (
	"encoding/json"
	"log"

	"github.com/ilgooz/mesg-go/service"
)

// Logger is a logger service.
type Logger struct {
	s *service.Service
}

// New creates a new Logger runs over service s.
func New(s *service.Service) *Logger {
	return &Logger{s: s}
}

// Start starts logger as a service.
func (l *Logger) Start() error {
	return l.s.Listen(
		service.NewTask("log", l.handler),
	)
}

func (l *Logger) handler(req *service.Request) service.Response {
	var data logRequest
	if err := req.Decode(&data); err != nil {
		return service.Response{
			"error": errorResponse{err.Error()},
		}
	}

	bytes, err := json.Marshal(data.Data)
	if err != nil {
		return service.Response{
			"error": errorResponse{err.Error()},
		}
	}

	log.Printf("%s: %s", data.ServiceID, string(bytes))

	return service.Response{
		"success": successResponse{"ok"},
	}
}

type logRequest struct {
	ServiceID string      `json:"serviceID"`
	Data      interface{} `json:"data"`
}

type successResponse struct {
	Message string `json:"message"`
}

type errorResponse struct {
	Message string `json:"message"`
}

// Close closes the service.
func (l *Logger) Close() error {
	return l.s.Close()
}
