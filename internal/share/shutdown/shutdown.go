package shutdown

import (
	"context"
	"errors"
	"log"
)

type Shutdown struct {
	Table map[string]func(string) error
	Ctx   context.Context
}

func Init(ctx context.Context) *Shutdown {
	return &Shutdown{
		Ctx:   ctx,
		Table: make(map[string]func(string) error),
	}
}

func (s *Shutdown) Wait() error {
	<-s.Ctx.Done()
	return s.Shutdown()
}

func (s *Shutdown) Shutdown() error {
	var totalErr error
	for id, service := range s.Table {
		log.Printf("Shutting down: %s", id)
		err := service(id)
		if err != nil {
			log.Printf("Error shutting down: %s | ERROR: %s", id, err)
			totalErr = errors.Join(totalErr, err)
		}
	}
	return totalErr
}

func (s *Shutdown) Push(funcid string, function func(string) error) {
	// WARN: May be throw error
	_, ok := s.Table[funcid]
	if ok {
		return
	}

	s.Table[funcid] = function
}

func (s *Shutdown) ShutdownByID(funcid string) error {
	function, ok := s.Table[funcid]
	if !ok {
		return errors.New("function does not exist")
	}
	return function(funcid)
}
