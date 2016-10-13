package gosync

import (
	"math/rand"
	"sync"
	"xnet/std/enc"
)

// A Map is a set of temporary objects that may be individually set, get and deleted.
//
// A Map is safe for use by multiple goroutines simultaneously.
type Map struct {
	mx   sync.RWMutex
	vals map[interface{}]interface{}
}

func normKey(key interface{}) interface{} {
	if v, ok := key.([]byte); ok {
		return string(v)
	}
	return key
}

func (m *Map) Clear() {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.vals = map[interface{}]interface{}{}
}

func (m *Map) Set(key, value interface{}) {
	m.mx.Lock()
	defer m.mx.Unlock()
	if m.vals == nil {
		m.vals = map[interface{}]interface{}{}
	}
	m.vals[normKey(key)] = value
}

func (m *Map) Delete(key interface{}) {
	m.mx.Lock()
	defer m.mx.Unlock()
	delete(m.vals, normKey(key))
}

func (m *Map) Get(key interface{}) interface{} {
	m.mx.RLock()
	defer m.mx.RUnlock()
	if v, ok := m.vals[normKey(key)]; ok {
		return v
	}
	return nil
}

func (m *Map) Exists(key interface{}) bool {
	m.mx.RLock()
	defer m.mx.RUnlock()
	_, ok := m.vals[normKey(key)]
	return ok
}

func (m *Map) Size() int {
	m.mx.RLock()
	defer m.mx.RUnlock()
	return len(m.vals)
}

func (m *Map) KeyValues() map[interface{}]interface{} {
	m.mx.RLock()
	defer m.mx.RUnlock()

	res := map[interface{}]interface{}{}
	for k, v := range m.vals {
		res[k] = v
	}
	return res
}

func (m *Map) Values() []interface{} {
	m.mx.RLock()
	defer m.mx.RUnlock()

	vv := make([]interface{}, 0, len(m.vals))
	for _, v := range m.vals {
		vv = append(vv, v)
	}
	return vv
}

func (m *Map) String() string {
	m.mx.RLock()
	defer m.mx.RUnlock()

	ss := map[string]string{}
	for k, v := range m.vals {
		ss[enc.String(k)] = enc.String(v)
	}
	return enc.String(ss)
}

func (m *Map) Strings() []string {
	ss := make([]string, 0, len(m.vals))
	for _, v := range m.Values() {
		ss = append(ss, enc.String(v))
	}
	return ss
}

func (m *Map) Pop() (key, value interface{}) {
	m.mx.Lock()
	defer m.mx.Unlock()
	for key, value = range m.vals {
		delete(m.vals, key)
		return
	}
	return
}

func (m *Map) RandomValue() interface{} {
	_, v := m.Random()
	return v
}

func (m *Map) Random() (key, value interface{}) {
	m.mx.RLock()
	defer m.mx.RUnlock()

	if cnt := len(m.vals); cnt > 0 {
		// todo: optimize it!  (add keys slice)
		n := rand.Intn(cnt)
		for k, v := range m.vals {
			if n == 0 {
				return k, v
			}
			n--
		}
		panic(1)
	}
	return nil, nil
}
