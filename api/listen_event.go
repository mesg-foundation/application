package api

import (
	"github.com/mesg-foundation/core/event"
	"github.com/mesg-foundation/core/service"
	"github.com/mesg-foundation/core/x/xstrings"
)

// EventListener provides functionalities to listen MESG events.
type EventListener struct {
	// Events receives matching events.
	Events chan *event.Event

	// Err filled when event subscription finished with a failure.
	Err chan error

	// cancel stops listening for new events.
	cancel chan struct{}

	// listening indicates if listening started
	listening chan struct{}

	// filters.
	eventKey string

	api *API
}

// ListenEventFilter is a filter func for filtering events.
type ListenEventFilter func(*EventListener)

// ListenEventKeyFilter returns an eventKey filter.
func ListenEventKeyFilter(eventKey string) ListenEventFilter {
	return func(ln *EventListener) {
		ln.eventKey = eventKey
	}
}

// ListenEvent listens events that matches with filters on service serviceID.
func (a *API) ListenEvent(serviceID string, filters ...ListenEventFilter) (*EventListener, error) {
	l := &EventListener{
		Events:    make(chan *event.Event),
		Err:       make(chan error, 1),
		cancel:    make(chan struct{}),
		listening: make(chan struct{}),
		api:       a,
	}
	for _, filter := range filters {
		filter(l)
	}
	return l, l.listen(serviceID)
}

// Close stops listening for events.
func (l *EventListener) Close() error {
	close(l.cancel)
	return nil
}

// listen listens events matches with eventFilter on serviceID.
func (l *EventListener) listen(serviceID string) error {
	s, err := l.api.db.Get(serviceID)
	if err != nil {
		return err
	}
	if err := l.validateEventKey(s); err != nil {
		return err
	}
	go l.listenLoop(s)
	<-l.listening
	return nil
}

func (l *EventListener) listenLoop(service *service.Service) {
	topic := service.EventSubscriptionChannel()
	subscription := l.api.ps.Sub(topic)
	defer l.api.ps.Unsub(subscription, topic)
	close(l.listening)

	for {
		select {
		case <-l.cancel:
			return

		// TODO use e.Err when subscription fails.
		// currently we don't need this but when pubsub refactored, we'll
		// need to pass an error to Err chan.
		case data := <-subscription:
			event := data.(*event.Event)
			if l.isSubscribedEvent(event) {
				l.Events <- event
			}
		}
	}
}

func (l *EventListener) validateEventKey(service *service.Service) error {
	if l.eventKey == "" || l.eventKey == "*" {
		return nil
	}
	_, err := service.GetEvent(l.eventKey)
	return err
}

func (l *EventListener) isSubscribedEvent(e *event.Event) bool {
	return xstrings.SliceContains([]string{"", "*", e.Key}, l.eventKey)
}
