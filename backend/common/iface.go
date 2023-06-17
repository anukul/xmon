package common

type DatabaseClient interface {
	GetScope() string
	Get(key string, result any) error
	Set(key string, value interface{}) error
	WithScope(scope string) DatabaseClient
}
