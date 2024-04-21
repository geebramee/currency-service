package cache

import (
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCacheGetSet(t *testing.T) {
	mockClient := redis.NewClient(&redis.Options{})

	c := &Cache{
		Client: mockClient,
	}

	key := "test_key"
	value := 10.0
	err := c.Set(key, value)
	assert.NoError(t, err)

	result, err := c.Get(key)
	assert.NoError(t, err)
	assert.Equal(t, value, result)
}
