package internal

type MutexMapPort interface {
	GetMutexMap(key string) (val []string, ok bool)
	SetMutexMap(key string, val []string)
}
