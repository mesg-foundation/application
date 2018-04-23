package pubsub

// Publish a message on a channel
func Publish(channel string, data Message) {
	mu.Lock()
	defer mu.Unlock()
	for _, listener := range listeners[channel] {
		if listener != nil {
			listener <- data
		}
	}
}
