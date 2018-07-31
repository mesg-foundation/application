package pubsub

import (
	"testing"

	"github.com/stvp/assert"
)

func TestPublish(t *testing.T) {
	key := "TestPublish"
	data := struct{}{}

	res := Subscribe(key)
	go Publish(key, data)
	x := <-res
	assert.Equal(t, x, data)
}

func TestPublishMultipleListeners(t *testing.T) {
	key := "TestPublishMultipleListeners"
	data := struct{}{}
	res1 := Subscribe(key)
	res2 := Subscribe(key)
	go Publish(key, data)
	x := <-res1
	y := <-res2
	assert.Equal(t, x, data)
	assert.Equal(t, y, data)
}

func TestSubscribe(t *testing.T) {
	key := "TestSubscribe"
	res := Subscribe(key)
	assert.NotNil(t, res)
	assert.Equal(t, len(listeners[key]), 1)
}

func TestSubscribeMultipleTimes(t *testing.T) {
	key := "TestSubscribeMultipleTimes"
	Subscribe(key)
	Subscribe(key)
	assert.Equal(t, len(listeners[key]), 2)
}
