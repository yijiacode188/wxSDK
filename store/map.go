package store

import (
	"sync"
	"time"
)

type ExpireValue struct {
	Value      any
	ExpireTime time.Time
}

var mapping sync.Map

func NewStorage() *Storage {
	return &Storage{}
}

type Storage struct {
}

func (s *Storage) SetStore(key string, value any, expire ...time.Time) error {
	data := &ExpireValue{
		Value: value,
	}
	if len(expire) != 0 {
		data.ExpireTime = expire[0]
	}
	mapping.Store(key, *data)
	return nil
}

func (s *Storage) GetStore(key string) (any, bool) {
	value, ok := mapping.Load(key)
	if !ok {
		return nil, false
	}
	isExpired := value.(ExpireValue).ExpireTime.Before(time.Now())
	if isExpired {
		//已过期
		s.DeleteStore(key)
		return nil, false
	} else {
		//没过期
		return value.(ExpireValue).Value, ok
	}
}

func (s *Storage) DeleteStore(key string) {
	mapping.Delete(key)
}
