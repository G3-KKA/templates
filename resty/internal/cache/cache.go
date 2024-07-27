package cache

type Cacher interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}
type cache struct {
}

func NewCacher() Cacher {
	return &cache{}
}

func (c *cache) Get(key string) (string, error) {
	return "", nil
}
func (c *cache) Set(key string, value string) error {
	return nil
}
