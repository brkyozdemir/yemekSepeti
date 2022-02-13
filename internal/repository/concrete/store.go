package concrete

type MemoryDatabase struct {
	Store []map[string]string
}

func NewMemDB() *MemoryDatabase {
	return &MemoryDatabase{
		Store: []map[string]string{},
	}
}

func (api *MemoryDatabase) GetObjectByKey(key string) map[string]string {
	for _, item := range api.Store {
		if _, ok := item[key]; ok {
			return item
		}
	}

	return map[string]string{"message": "Key not found!"}
}

func (api *MemoryDatabase) GetObjectList() []map[string]string {
	return api.Store
}

func (api *MemoryDatabase) CreateObject(key string, value string) map[string]string {
	for _, item := range api.Store {
		if _, ok := item[key]; ok {
			return item
		}
	}

	store := append(api.Store, map[string]string{key: value})

	api.Store = store
	return map[string]string{key: value}
}

func (api *MemoryDatabase) FlushObjectList() (bool, error) {
	api.Store = []map[string]string{}

	if len(api.Store) > 0 {
		return false, nil
	}

	return true, nil
}
