package core

import (
	"encoding/json"

	"github.com/mesg-foundation/core/api"
	"github.com/mesg-foundation/core/protobuf/acknowledgement"
	"github.com/mesg-foundation/core/protobuf/coreapi"
)

// ListenEvent listens events matches with eventFilter on serviceID.
func (s *Server) ListenEvent(request *coreapi.ListenEventRequest, stream coreapi.Core_ListenEventServer) error {
	ln, err := s.api.ListenEvent(request.ServiceID, api.ListenEventKeyFilter(request.EventFilter))
	if err != nil {
		return err
	}
	defer ln.Close()

	// send header to notify client that the stream is ready.
	if err := acknowledgement.SetStreamReady(stream); err != nil {
		return err
	}

	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case err := <-ln.Err:
			return err

		case ev := <-ln.Events:
			evData, err := json.Marshal(ev.Data)
			if err != nil {
				return err
			}

			if err := stream.Send(&coreapi.EventData{
				EventKey:  ev.Key,
				EventData: string(evData),
			}); err != nil {
				return err
			}
		}
	}
}
