package service

import "sync"

type StatusService struct {
	mu    sync.RWMutex
	ready bool
}

func NewStatusService() *StatusService {
	return &StatusService{ready: true}
}

func (s *StatusService) SetReady(ready bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ready = ready
}

func (s *StatusService) IsReady() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.ready
}
