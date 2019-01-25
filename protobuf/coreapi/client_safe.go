package coreapi

import (
	"context"
	"io"
	"time"

	"google.golang.org/grpc"
)

const reconnectDelay = 3 * time.Second

// CoreClientSafe provides CoreClient with stream reconnection.
type CoreClientSafe struct {
	CoreClient
}

// NewCoreClient creates core client with reconnection.
func NewCoreClientSafe(cc *grpc.ClientConn) *CoreClientSafe {
	return &CoreClientSafe{
		CoreClient: NewCoreClient(cc),
	}
}

// core_ListenEventClient is a client with reconnection.
type core_ListenEventClient struct {
	Core_ListenEventClient

	client *CoreClientSafe
	c      chan *core_ListenEventClientResponse

	ctx  context.Context
	in   *ListenEventRequest
	opts []grpc.CallOption
}

// core_ListenEventClientResponse wraps ListenEvent recv response.
type core_ListenEventClientResponse struct {
	eventData *EventData
	err       error
}

// newCore_ListenEventClient creates core ListenEvent client.
func newCore_ListenEventClient(client *CoreClientSafe, ctx context.Context, in *ListenEventRequest, opts ...grpc.CallOption) *core_ListenEventClient {
	c := &core_ListenEventClient{
		client: client,
		c:      make(chan *core_ListenEventClientResponse),

		ctx:  ctx,
		in:   in,
		opts: opts,
	}
	waitStream := make(chan struct{}, 1)
	go c.recvLoop(waitStream)
	<-waitStream
	return c
}

// recvLoop receives ListenEvent response in loop and reconnect in on error.
func (s *core_ListenEventClient) recvLoop(waitStream chan struct{}) {
	var err error
loop:
	for {
		// connect
		s.Core_ListenEventClient, err = s.client.CoreClient.ListenEvent(s.ctx, s.in, s.opts...)
		waitStream <- struct{}{}
		if err != nil {
			s.c <- &core_ListenEventClientResponse{nil, err}
			continue
		}

		// buffered channel, because it might happen that ctx.Done
		// will be notified first before stream.Recv in for loop.
		done := make(chan struct{}, 1)

		go func(c Core_ListenEventClient) {
			select {
			case <-c.Context().Done():
				c.CloseSend()
			case <-done:
			}
		}(s.Core_ListenEventClient)

		for {
			td, err := s.Core_ListenEventClient.Recv()
			s.c <- &core_ListenEventClientResponse{td, err}
			if err != nil {
				done <- struct{}{}
				// in case of EOF end loop
				if err == io.EOF {
					break loop
				}
				break
			}
		}

		// sleep before reconnect
		time.Sleep(reconnectDelay)
	}
}

// Recv receives data from streams.
func (s *core_ListenEventClient) Recv() (*EventData, error) {
	v := <-s.c
	return v.eventData, v.err
}

// ListenEvent subscribe to a stream that listens for events from a service.
func (c *CoreClientSafe) ListenEvent(ctx context.Context, in *ListenEventRequest, opts ...grpc.CallOption) (Core_ListenEventClient, error) {
	return newCore_ListenEventClient(c, ctx, in, opts...), nil
}

// core_ListenResultClient is a client with reconnection.
type core_ListenResultClient struct {
	Core_ListenResultClient

	client *CoreClientSafe
	c      chan *core_ListenResultClientResponse

	ctx  context.Context
	in   *ListenResultRequest
	opts []grpc.CallOption
}

// core_ListenResultClientResponse wraps ListenResult recv response.
type core_ListenResultClientResponse struct {
	resultData *ResultData
	err        error
}

// newCore_ListenResultClient creates core ListenResult client.
func newCore_ListenResultClient(client *CoreClientSafe, ctx context.Context, in *ListenResultRequest, opts ...grpc.CallOption) *core_ListenResultClient {
	c := &core_ListenResultClient{
		client: client,
		c:      make(chan *core_ListenResultClientResponse),

		ctx:  ctx,
		in:   in,
		opts: opts,
	}
	waitStream := make(chan struct{}, 1)
	go c.recvLoop(waitStream)
	<-waitStream
	return c
}

// recvLoop receives ListenResult response in loop and reconnect in on error.
func (s *core_ListenResultClient) recvLoop(waitStream chan struct{}) {
	var err error
loop:
	for {
		// connect
		s.Core_ListenResultClient, err = s.client.CoreClient.ListenResult(s.ctx, s.in, s.opts...)
		waitStream <- struct{}{}
		if err != nil {
			s.c <- &core_ListenResultClientResponse{nil, err}
			continue
		}

		// buffered channel, because it might happen that ctx.Done
		// will be notified first before stream.Recv in for loop.
		done := make(chan struct{}, 1)

		go func(c Core_ListenResultClient) {
			select {
			case <-c.Context().Done():
				c.CloseSend()
			case <-done:
			}
		}(s.Core_ListenResultClient)

		for {
			td, err := s.Core_ListenResultClient.Recv()
			s.c <- &core_ListenResultClientResponse{td, err}
			if err != nil {
				done <- struct{}{}
				// in case of EOF end loop
				if err == io.EOF {
					break loop
				}
				break
			}
		}

		// sleep before reconnect
		time.Sleep(reconnectDelay)
	}
}

// Recv receives data from streams.
func (s *core_ListenResultClient) Recv() (*ResultData, error) {
	v := <-s.c
	return v.resultData, v.err
}

func (c *CoreClientSafe) ListenResult(ctx context.Context, in *ListenResultRequest, opts ...grpc.CallOption) (Core_ListenResultClient, error) {
	return newCore_ListenResultClient(c, ctx, in, opts...), nil
}

// core_ServiceLogsClient is a client with reconnection.
type core_ServiceLogsClient struct {
	Core_ServiceLogsClient

	client *CoreClientSafe
	c      chan *core_ServiceLogsClientResponse

	ctx  context.Context
	in   *ServiceLogsRequest
	opts []grpc.CallOption
}

// core_ServiceLogsClientResponse wraps ServiceLogs recv response.
type core_ServiceLogsClientResponse struct {
	logData *LogData
	err     error
}

// newCore_ServiceLogsClient creates core ServiceLogs client.
func newCore_ServiceLogsClient(client *CoreClientSafe, ctx context.Context, in *ServiceLogsRequest, opts ...grpc.CallOption) *core_ServiceLogsClient {
	c := &core_ServiceLogsClient{
		client: client,
		c:      make(chan *core_ServiceLogsClientResponse),

		ctx:  ctx,
		in:   in,
		opts: opts,
	}
	waitStream := make(chan struct{}, 1)
	go c.recvLoop(waitStream)
	<-waitStream
	return c
}

// recvLoop receives ServiceLogs response in loop and reconnect in on error.
func (s *core_ServiceLogsClient) recvLoop(waitStream chan struct{}) {
	var err error
loop:
	for {
		// connect
		s.Core_ServiceLogsClient, err = s.client.CoreClient.ServiceLogs(s.ctx, s.in, s.opts...)
		waitStream <- struct{}{}
		if err != nil {
			s.c <- &core_ServiceLogsClientResponse{nil, err}
			continue
		}

		// buffered channel, because it might happen that ctx.Done
		// will be notified first before stream.Recv in for loop.
		done := make(chan struct{}, 1)

		go func(c Core_ServiceLogsClient) {
			select {
			case <-c.Context().Done():
				c.CloseSend()
			case <-done:
			}
		}(s.Core_ServiceLogsClient)

		for {
			td, err := s.Core_ServiceLogsClient.Recv()
			s.c <- &core_ServiceLogsClientResponse{td, err}
			if err != nil {
				done <- struct{}{}
				// in case of EOF end loop
				if err == io.EOF {
					break loop
				}
				break
			}
		}

		// sleep before reconnect
		time.Sleep(reconnectDelay)
	}
}

// Recv receives data from streams.
func (s *core_ServiceLogsClient) Recv() (*LogData, error) {
	v := <-s.c
	return v.logData, v.err
}

func (c *CoreClientSafe) ServiceLogs(ctx context.Context, in *ServiceLogsRequest, opts ...grpc.CallOption) (Core_ServiceLogsClient, error) {
	return newCore_ServiceLogsClient(c, ctx, in, opts...), nil
}
