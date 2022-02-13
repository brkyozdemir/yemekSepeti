package repository

type StoreDatabase interface {
	GetObjectByKey(key string) map[string]string
	GetObjectList() []map[string]string
	CreateObject(key string, value string) map[string]string
	FlushObjectList() (bool, error)
}
