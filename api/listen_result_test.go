package api

import (
	"testing"

	"github.com/mesg-foundation/core/execution"
	"github.com/mesg-foundation/core/service"
	"github.com/stretchr/testify/require"
)

func TestValidateTaskKey(t *testing.T) {
	var (
		l = &ResultListener{}
		s = &service.Service{
			Tasks: []*service.Task{
				{
					Key: "test",
				},
			},
		}
	)

	l.taskKey = ""
	require.Nil(t, l.validateTaskKey(s))

	l.taskKey = "*"
	require.Nil(t, l.validateTaskKey(s))

	l.taskKey = "test"
	require.Nil(t, l.validateTaskKey(s))

	l.taskKey = "xxx"
	require.NotNil(t, l.validateTaskKey(s))
}

func TestValidateOutputKey(t *testing.T) {
	var (
		l = &ResultListener{}
		s = &service.Service{
			Tasks: []*service.Task{
				{
					Key: "test",
					Outputs: []*service.Output{
						{
							Key: "outputx",
						},
					},
				},
			},
		}
	)

	l.taskKey = "test"
	l.outputKey = ""
	require.Nil(t, l.validateOutputKey(s))

	l.taskKey = "test"
	l.outputKey = "*"
	require.Nil(t, l.validateOutputKey(s))

	l.taskKey = "test"
	l.outputKey = "outputx"
	require.Nil(t, l.validateOutputKey(s))

	l.taskKey = "test"
	l.outputKey = "xxx"
	require.NotNil(t, l.validateOutputKey(s))

	l.taskKey = "xxx"
	l.outputKey = ""
	require.Nil(t, l.validateOutputKey(s))

	l.taskKey = "xxx"
	l.outputKey = "*"
	require.Nil(t, l.validateOutputKey(s))

	l.taskKey = "xxx"
	l.outputKey = "outputX"
	require.NotNil(t, l.validateOutputKey(s))

	l.taskKey = "xxx"
	l.outputKey = "xxx"
	require.NotNil(t, l.validateOutputKey(s))
}

func TestIsSubscribedToTask(t *testing.T) {
	var (
		l = &ResultListener{}
		x = &execution.Execution{TaskKey: "task"}
	)

	l.taskKey = ""
	require.True(t, l.isSubscribedToTask(x))

	l.taskKey = "*"
	require.True(t, l.isSubscribedToTask(x))

	l.taskKey = "task"
	require.True(t, l.isSubscribedToTask(x))

	l.taskKey = "xxx"
	require.False(t, l.isSubscribedToTask(x))
}

func TestIsSubscribedToOutput(t *testing.T) {
	var (
		l = &ResultListener{}
		x = &execution.Execution{OutputKey: "output"}
	)

	l.outputKey = ""
	require.True(t, l.isSubscribedToOutput(x))

	l.outputKey = "*"
	require.True(t, l.isSubscribedToOutput(x))

	l.outputKey = "output"
	require.True(t, l.isSubscribedToOutput(x))

	l.outputKey = "xxx"
	require.False(t, l.isSubscribedToOutput(x))
}

func TestIsSubscribedToTags(t *testing.T) {
	l := &ResultListener{}

	type result struct {
		execution *execution.Execution
		valid     bool
	}
	tests := []struct {
		tags    []string
		results []result
	}{
		{
			[]string{},
			[]result{
				{&execution.Execution{}, true},
				{&execution.Execution{Tags: []string{"foo"}}, true},
				{&execution.Execution{Tags: []string{"foo", "bar"}}, true},
				{&execution.Execution{Tags: []string{"none"}}, true},
			},
		},
		{
			[]string{"foo"},
			[]result{
				{&execution.Execution{}, false},
				{&execution.Execution{Tags: []string{"foo"}}, true},
				{&execution.Execution{Tags: []string{"foo", "bar"}}, true},
				{&execution.Execution{Tags: []string{"none"}}, false},
			},
		},
		{
			[]string{"foo", "bar"},
			[]result{
				{&execution.Execution{}, false},
				{&execution.Execution{Tags: []string{"foo"}}, false},
				{&execution.Execution{Tags: []string{"foo", "bar"}}, true},
				{&execution.Execution{Tags: []string{"none"}}, false},
			},
		},
	}
	for _, test := range tests {
		for _, r := range test.results {
			l.tagFilters = test.tags
			require.Equal(t, r.valid, l.isSubscribedToTags(r.execution))
		}
	}
}

func TestIsSubscribed(t *testing.T) {
	l := &ResultListener{}

	type test struct {
		taskFilter, outputFilter string
		tagFilters               []string
		e                        execution.Execution
	}
	subscribeToOutput := func(x *test) *test {
		return x
	}
	notSubscribeToOutput := func(x *test) *test {
		x.outputFilter = "foo"
		return x
	}
	subscribeToTask := func(x *test) *test {
		return x
	}
	notSubscribeToTask := func(x *test) *test {
		x.taskFilter = "foo"
		return x
	}
	subscribeToTags := func(x *test) *test {
		return x
	}
	notSubscribeToTags := func(x *test) *test {
		x.tagFilters = []string{"foo"}
		return x
	}
	tests := []struct {
		t     *test
		valid bool
		msg   string
	}{
		{notSubscribeToTags(notSubscribeToTask(notSubscribeToOutput(&test{}))), false, "[]"},
		{notSubscribeToTags(notSubscribeToTask(subscribeToOutput(&test{}))), false, "[output]"},
		{notSubscribeToTags(subscribeToTask(notSubscribeToOutput(&test{}))), false, "[task]"},
		{notSubscribeToTags(subscribeToTask(subscribeToOutput(&test{}))), false, "[task, output]"},
		{subscribeToTags(notSubscribeToTask(notSubscribeToOutput(&test{}))), false, "[tags]"},
		{subscribeToTags(notSubscribeToTask(subscribeToOutput(&test{}))), false, "[tags, output]"},
		{subscribeToTags(subscribeToTask(notSubscribeToOutput(&test{}))), false, "[tags, task]"},
		{subscribeToTags(subscribeToTask(subscribeToOutput(&test{}))), true, "[tags, task, output]"},
	}
	for _, test := range tests {
		l.taskKey = test.t.taskFilter
		l.outputKey = test.t.outputFilter
		l.tagFilters = test.t.tagFilters
		require.Equal(t, test.valid, l.isSubscribed(&test.t.e), test.msg)
	}
}