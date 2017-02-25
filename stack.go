package gosync

import "sync"

type Stack struct {
	mx   sync.RWMutex
	vals []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.vals = append(s.vals, value)
}

func (s *Stack) Pop() interface{} {
	s.mx.Lock()
	defer s.mx.Unlock()
	if len(s.vals) > 0 {
		v := s.vals[len(s.vals)-1]
		s.vals = s.vals[:len(s.vals)-1]
		return v
	}
	return nil
}

func (s *Stack) Size() int {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return len(s.vals)
}

func (s *Stack) Values() []interface{} {
	s.mx.RLock()
	defer s.mx.RUnlock()

	vv := make([]interface{}, len(s.vals))
	copy(vv, s.vals)
	return vv
}

func (s *Stack) Clear() {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.vals = nil
}
