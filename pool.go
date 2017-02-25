package gosync

import (
	"sync"
	"xnet/std/enc"
)

type Pool struct {
	mx   sync.RWMutex
	vals []interface{}
}

func (q *Pool) Push(value interface{}) {
	q.mx.Lock()
	defer q.mx.Unlock()
	q.vals = append(q.vals, value)
}

func (q *Pool) Pop() interface{} {
	q.mx.Lock()
	defer q.mx.Unlock()
	if len(q.vals) > 0 {
		v := q.vals[0]
		q.vals = q.vals[1:]
		return v
	}
	return nil
}

func (q *Pool) Size() int {
	q.mx.RLock()
	defer q.mx.RUnlock()
	return len(q.vals)
}

func (q *Pool) Values() []interface{} {
	q.mx.RLock()
	defer q.mx.RUnlock()

	vv := make([]interface{}, len(q.vals))
	copy(vv, q.vals)
	return vv
}

func (q *Pool) String() string {
	q.mx.RLock()
	defer q.mx.RUnlock()

	return enc.String(q.vals)
}

func (q *Pool) Strings() []string {
	q.mx.RLock()
	defer q.mx.RUnlock()

	ss := make([]string, 0, len(q.vals))
	for _, v := range q.vals {
		ss = append(ss, enc.String(v))
	}
	return ss
}

func (q *Pool) Clear() {
	q.mx.Lock()
	defer q.mx.Unlock()
	q.vals = nil
}
