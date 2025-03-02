package utils

import (
	"sync"
	"time"
)

// SingleValueCache implements a simple cache that stores a single value
// with optional expiration.
type SingleValueCache[T any] struct {
	mu        sync.RWMutex
	value     T
	expiresAt time.Time
	hasValue  bool
}

// NewSingleValueCache creates a new instance of a single value cache.
func NewSingleValueCache[T any]() *SingleValueCache[T] {
	return &SingleValueCache[T]{}
}

// SetWithExpiry stores a value in the cache with an expiration duration.
func (c *SingleValueCache[T]) SetWithExpiry(value T, expiry time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value = value
	c.hasValue = true
	c.expiresAt = time.Now().Add(expiry)
}

// Get retrieves the stored value from the cache.
// Returns the value and a boolean indicating if the value was found and not expired.
func (c *SingleValueCache[T]) Get() (T, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var zero T

	if !c.HasValue() {
		return zero, false
	}

	return c.value, true
}

// Clear removes the value from the cache.
func (c *SingleValueCache[T]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.hasValue = false

	var zero T
	c.value = zero
}

// HasValue returns whether the cache currently has a valid value.
func (c *SingleValueCache[T]) HasValue() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if !c.hasValue || time.Now().After(c.expiresAt) {
		return false
	}

	return true
}

// TimeToExpiry returns the duration until expiration.
// If the value has no expiry or is already expired, returns -1.
func (c *SingleValueCache[T]) TimeToExpiry() time.Duration {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if !c.HasValue() {
		return -1
	}

	remaining := time.Until(c.expiresAt)

	return remaining
}
