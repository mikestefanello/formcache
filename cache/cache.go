package cache

import (
	"errors"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Cache struct {
	Rows map[string]CacheData
}

type CacheData struct {
	ID        string
	Data      map[string]interface{}
	Expires   time.Time
	CreatedAt time.Time
}

var ErrNotFound = errors.New("Cache entry not found")

func NewCache() *Cache {
	return &Cache{Rows: make(map[string]CacheData)}
}

func NewCacheData() CacheData {
	return CacheData{Data: make(map[string]interface{})}
}

func (c *Cache) GetData(id string) (CacheData, error) {
	d, ok := c.Rows[id]
	if !ok {
		return d, ErrNotFound
	}
	return d, nil
}

func (c *Cache) SetData(d *CacheData) error {
	if d.ID == "" {
		d.ID = uuid.NewV4().String()
	}
	d.CreatedAt = time.Now()
	c.Rows[d.ID] = *d
	return nil
}

func (c *Cache) DeleteData(id string) error {
	delete(c.Rows, id)
	return nil
}

func (c *Cache) GetFormCache(r *http.Request) (CacheData, error) {
	if r.FormValue("token") != "" {
		return c.GetData(r.FormValue("token"))
	}
	return CacheData{}, ErrNotFound
}
