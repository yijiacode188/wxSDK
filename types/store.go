package types

import "time"

type StoreInterface interface {
	SetStore(key string, value any, expire ...time.Time) error
	GetStore(key string) (any, bool)
	DeleteStore(key string)
}
