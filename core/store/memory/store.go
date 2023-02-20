package memory

import (
	"errors"
	"fastserver/core/util"
	"sync"
	"time"
)

// 内存数据存储
type Store struct {
	sync.RWMutex
	values map[string]*Record
}

func (ms *Store) List() ([]*Record, error) {
	ms.RLock()
	defer ms.RUnlock()

	var values []*Record

	for _, v := range ms.values {
		if v.CheckState() {
			values = append(values, v)
		}
	}

	return values, nil
}

func (ms *Store) Read(keys ...string) ([]*Record, error) {
	ms.RLock()
	defer ms.RUnlock()

	var records []*Record

	for _, key := range keys {
		v, ok := ms.values[key]
		if !ok {
			return nil, errors.New("NotFound")
		}

		if !v.CheckState() {
			return nil, errors.New("NotFound")
		}

		records = append(records, v)
	}

	return records, nil
}

func (ms *Store) Write(records ...*Record) error {
	ms.Lock()
	defer ms.Unlock()

	for _, r := range records {
		ms.values[r.key] = r
	}

	return nil
}

func (ms *Store) Delete(keys ...string) error {
	ms.Lock()
	defer ms.Unlock()

	for _, key := range keys {
		delete(ms.values, key)
	}

	return nil
}

func (ms *Store) Get(key string) (interface{}, error) {
	ms.RLock()
	defer ms.RUnlock()

	v, ok := ms.values[key]
	if !ok {
		return nil, errors.New("NotFound")
	}

	return v.Value(), nil
}

func (ms *Store) Set(key string, value interface{}, expiry ...time.Duration) error {
	return ms.Write(NewRecord(key, value, expiry...))
}

func (ms *Store) Int(key string) (int, error) {
	val, err := ms.Get(key)
	if err != nil {
		return 0, err
	}
	return util.ParseInt(val), nil
}

func (ms *Store) Int32(key string) (int32, error) {
	val, err := ms.Get(key)
	if err != nil {
		return int32(0), err
	}
	return util.ParseInt32(val), nil
}

func (ms *Store) Int64(key string) (int64, error) {
	val, err := ms.Get(key)
	if err != nil {
		return 0, err
	}
	return util.ParseInt64(val), nil
}

func (ms *Store) Uint32(key string) (uint32, error) {
	val, err := ms.Get(key)
	if err != nil {
		return 0, err
	}
	return util.ParseUint32(val), nil
}

func (ms *Store) Uint64(key string) (uint64, error) {
	val, err := ms.Get(key)
	if err != nil {
		return 0, err
	}
	return util.ParseUint64(val), nil
}

func (ms *Store) Bool(key string) (bool, error) {
	val, err := ms.Get(key)
	if err != nil {
		return false, err
	}
	return util.ParseBool(val), nil
}

// NewStore returns a new store.Store
func NewStore() *Store {
	return &Store{
		values: make(map[string]*Record),
	}
}
