package sctructs

import (
	"sync"
	"time"
)

// AuthToken - struct for tokens to auth
type AuthToken struct {
	// user info
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	// token itself
	Value string `json:"value"`

	// expiration time
	TimeCreated time.Time `json:"time_created"`
	// time expired
	TimeExpired time.Time `json:"time_expired"`
}

func (t *AuthToken) Expired() bool {
	return time.Now().After(t.TimeExpired)
}

// AuthTokenCache - cache with tokens
type AuthTokenCache struct {
	m   map[string]AuthToken
	mut sync.Mutex
}

// Init - returns valid struct
func (c AuthTokenCache) Init() AuthTokenCache {
	return AuthTokenCache{
		m:   make(map[string]AuthToken),
		mut: sync.Mutex{},
	}
}

func (c *AuthTokenCache) Get(key string) (AuthToken, bool) {
	c.mut.Lock()
	result, ok := c.m[key]
	c.mut.Unlock()
	return result, ok
}

func (c *AuthTokenCache) Insert(key string, value AuthToken) {
	c.mut.Lock()
	c.m[key] = value
	c.mut.Unlock()
}

func (c *AuthTokenCache) Delete(key string) {
	c.mut.Lock()
	delete(c.m, key)
	c.mut.Unlock()
}
