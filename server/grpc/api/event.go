package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/mesg-foundation/engine/protobuf/acknowledgement"
	"github.com/mesg-foundation/engine/protobuf/api"
	"github.com/mesg-foundation/engine/sdk"
	eventsdk "github.com/mesg-foundation/engine/sdk/event"
)

// EventServer serve event functions.
type EventServer struct {
	sdk *sdk.SDK
}

// NewEventServer creates a new EventServer.
func NewEventServer(sdk *sdk.SDK) *EventServer {
	return &EventServer{sdk: sdk}
}

// Create creates a new event.
func (s *EventServer) Create(ctx context.Context, req *api.CreateEventRequest) (*api.CreateEventResponse, error) {
	if req.Key == "" {
		return nil, errors.New("create event: key missing")
	}

	event, err := s.sdk.Event.Create(req.InstanceHash, req.Key, req.Data)
	if err != nil {
		return nil, fmt.Errorf("create event: data %s", err)
	}

	return &api.CreateEventResponse{Hash: event.Hash}, nil
}

// Stream returns stream of events.
func (s *EventServer) Stream(req *api.StreamEventRequest, resp api.Event_StreamServer) error {
	var f *eventsdk.Filter
	if req.Filter != nil {
		f = &eventsdk.Filter{
			Hash:         req.Filter.Hash,
			InstanceHash: req.Filter.InstanceHash,
			Key:          req.Filter.Key,
		}
	}
	stream := s.sdk.Event.GetStream(f)
	defer stream.Close()

	// send header to notify client that the stream is ready.
	if err := acknowledgement.SetStreamReady(resp); err != nil {
		return err
	}

	for event := range stream.C {
		if err := resp.Send(e); err != nil {
			return err
		}
	}

	return nil
}
