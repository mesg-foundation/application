package service

import (
	"io"
	"sync"
	"testing"

	"github.com/docker/docker/pkg/stdcopy"
	"github.com/mesg-foundation/core/container"
	"github.com/mesg-foundation/core/container/mocks"
	"github.com/stretchr/testify/require"
)

func TestDependencyLogs(t *testing.T) {
	testDependencyLogs(t, func(s *Service, c container.Container, dependencyKey string) (rstd, rerr io.ReadCloser,
		err error) {
		dep, err := s.getDependency(dependencyKey)
		require.NoError(t, err)
		return dep.Logs(c, s.namespace())
	})
}

func testDependencyLogs(t *testing.T, do func(s *Service, c container.Container, dependencyKey string) (rstd, rerr io.ReadCloser,
	err error)) {
	var (
		dependencyKey = "1"
		stdData       = []byte{1, 2}
		errData       = []byte{3, 4}
		s             = &Service{
			Dependencies: []*Dependency{
				{Key: dependencyKey},
			},
		}
		mc = &mocks.Container{}
	)

	rp, wp := io.Pipe()
	wstd := stdcopy.NewStdWriter(wp, stdcopy.Stdout)
	werr := stdcopy.NewStdWriter(wp, stdcopy.Stderr)

	go wstd.Write(stdData)
	go werr.Write(errData)

	d, _ := s.getDependency(dependencyKey)
	mc.On("ServiceLogs", d.namespace(s.namespace())).Once().Return(rp, nil)

	rstd, rerr, err := do(s, mc, dependencyKey)
	require.NoError(t, err)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		buf := make([]byte, 2)
		_, err := rstd.Read(buf)
		require.NoError(t, err)
		require.Equal(t, stdData, buf)
	}()

	go func() {
		defer wg.Done()
		buf := make([]byte, 2)
		_, err = rerr.Read(buf)
		require.NoError(t, err)
		require.Equal(t, errData, buf)
	}()

	wg.Wait()
	mc.AssertExpectations(t)
}
