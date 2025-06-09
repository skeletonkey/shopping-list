package eventing

func TouchFileExists() (bool, error) {
	cfg := getConfig()

	age, err := cfg.touchFileAge()

	return age > 0, err
}

var e events
var eInit bool = false

func InitEventing() (events, error) {
	var err error
	if !eInit {
		e = events{
			cache:        make(cacheEntries, 0),
			cacheUpdated: false,
		}

		err = e.loadCache()
	}

	return e, err
}

func (e events) loadCache() error {
	e.cache[getCacheKey(1, 2)] = true
	e.cache[getCacheKey(3, 4)] = false

	return nil
}

func (e events) AddEntry(listId, itemId int) {
	e.cache[getCacheKey(listId, itemId)] = true
	e.cacheUpdated = true
}

func (e events) RemoveEntry(listId, itemId int) {
	e.cache[getCacheKey(listId, itemId)] = false
	e.cacheUpdated = true
}

func getCacheKey(listId, itemId int) cache {
	return cache{listId, itemId}
}
