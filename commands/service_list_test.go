package commands

import (
	"bufio"
	"io"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mesg-foundation/core/protobuf/coreapi"
)

func TestServiceList(t *testing.T) {
	var (
		services = []*coreapi.Service{
			{ID: "1", Name: "a", Status: coreapi.Service_RUNNING},
			{ID: "2", Name: "b", Status: coreapi.Service_PARTIAL},
		}
		m    = &mockServiceExecutor{}
		r, w = io.Pipe()
		tw   = &testTableWriter{w}
		br   = bufio.NewReader(r)
	)

	c := newServiceListCmd(m, tw)
	m.On("ServiceList").Return(services, nil)
	go c.runE(nil, nil)

	matched, err := regexp.Match(`\s*^STATUS\s+SERVICE\s+NAME\s*$`, readLine(t, br))
	require.NoError(t, err)
	require.True(t, matched)

	matched, err = regexp.Match(`\s*^running\s+1\s+a`, readLine(t, br))
	require.NoError(t, err)
	require.True(t, matched)

	matched, err = regexp.Match(`\s*^partial\s+2\s+b`, readLine(t, br))
	require.NoError(t, err)
	require.True(t, matched)
}

type testTableWriter struct {
	w io.Writer
}

func (t *testTableWriter) Write(b []byte) (n int, err error) {
	return t.w.Write(b)
}

func readLine(t *testing.T, r *bufio.Reader) []byte {
	line, _, err := r.ReadLine()
	require.NoError(t, err)
	return line
}
