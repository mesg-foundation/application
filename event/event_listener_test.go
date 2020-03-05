package event

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cskr/pubsub"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/crypto"
)

func TestFilter(t *testing.T) {
	var tests = []struct {
		f     *Filter
		e     *Event
		match bool
	}{
		{
			nil,
			nil,
			true,
		},
		{
			&Filter{},
			&Event{},
			true,
		},
		{
			&Filter{Hash: sdk.AccAddress(crypto.AddressHash([]byte("1")))},
			&Event{Hash: sdk.AccAddress(crypto.AddressHash([]byte("1")))},
			true,
		},
		{
			&Filter{Hash: sdk.AccAddress(crypto.AddressHash([]byte("1")))},
			&Event{Hash: sdk.AccAddress(crypto.AddressHash([]byte("2")))},
			false,
		},
		{
			&Filter{InstanceHash: sdk.AccAddress(crypto.AddressHash([]byte("1")))},
			&Event{InstanceHash: sdk.AccAddress(crypto.AddressHash([]byte("1")))},
			true,
		},
		{
			&Filter{InstanceHash: sdk.AccAddress(crypto.AddressHash([]byte("1")))},
			&Event{InstanceHash: sdk.AccAddress(crypto.AddressHash([]byte("1")))},
			true,
		},
		{
			&Filter{Key: "0"},
			&Event{Key: "0"},
			true,
		},
		{
			&Filter{Key: "*"},
			&Event{Key: "0"},
			true,
		},
		{
			&Filter{Key: "0"},
			&Event{Key: "1"},
			false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.match, tt.f.Match(tt.e))
	}
}

func TestEventListener(t *testing.T) {
	topic := "test-topic"
	testEvent := &Event{Key: "0"}
	ps := pubsub.New(0)
	el := NewListener(ps, topic, &Filter{Key: "0"})

	go func() {
		ps.Pub(&Event{Key: "1"}, topic)
		ps.Pub(testEvent, topic)
	}()
	go el.Listen()

	recvEvent := <-el.C
	assert.Equal(t, testEvent, recvEvent)

	el.Close()
	_, ok := <-el.C
	assert.False(t, ok)
}
