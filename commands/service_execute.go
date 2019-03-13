package commands

import (
	"encoding/json"
	"fmt"

	"github.com/mesg-foundation/core/protobuf/coreapi"
	"github.com/mesg-foundation/core/utils/pretty"
	casting "github.com/mesg-foundation/core/utils/servicecasting"
	"github.com/mesg-foundation/core/x/xjson"
	"github.com/mesg-foundation/core/x/xpflag"
	"github.com/mesg-foundation/core/x/xstrings"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

type serviceExecuteCmd struct {
	baseCmd

	executeData map[string]string
	taskKey     string
	jsonFile    string

	e ServiceExecutor
}

func newServiceExecuteCmd(e ServiceExecutor) *serviceExecuteCmd {
	c := &serviceExecuteCmd{e: e}
	c.cmd = newCommand(&cobra.Command{
		Use:     "execute",
		Short:   "Execute a task of a service",
		Example: `mesg-core service execute SERVICE`,
		Args:    cobra.ExactArgs(1),
		PreRunE: c.preRunE,
		RunE:    c.runE,
	})
	c.cmd.Flags().StringVarP(&c.taskKey, "task", "t", c.taskKey, "Run the given task")
	c.cmd.Flags().VarP(xpflag.NewStringToStringValue(&c.executeData, nil), "data", "d", "data required to run the task")
	c.cmd.Flags().StringVarP(&c.jsonFile, "json", "j", c.jsonFile, "Path to a JSON file containing the data required to run the task")
	return c
}

func (c *serviceExecuteCmd) preRunE(cmd *cobra.Command, args []string) error {
	if cmd.Flag("data").Changed && cmd.Flag("json").Changed {
		return errors.New("only one of '--data' or '--json' options can be specified")
	}
	return nil
}

func (c *serviceExecuteCmd) runE(cmd *cobra.Command, args []string) error {
	var (
		s              *coreapi.Service
		result         *coreapi.ResultData
		listenResultsC chan *coreapi.ResultData
		inputData      string
		resultsErrC    chan error
		err            error
	)

	pretty.Progress("Loading the service...", func() {
		s, err = c.e.ServiceByID(args[0])
	})
	if err != nil {
		return err
	}

	if err = c.getTaskKey(s); err != nil {
		return err
	}

	inputData, err = c.getData(s)
	if err != nil {
		return err
	}
	pretty.Progress(fmt.Sprintf("Executing task %q...", c.taskKey), func() {
		// Create an unique tag that will be used to listen to the result of this exact execution
		tags := []string{uuid.NewV4().String()}

		listenResultsC, resultsErrC, err = c.e.ServiceListenResults(args[0], c.taskKey, "", tags)
		if err != nil {
			return
		}

		err = c.e.ServiceExecuteTask(args[0], c.taskKey, inputData, tags)
	})
	if err != nil {
		return err
	}
	fmt.Printf("%s Task %q executed\n", pretty.SuccessSign, c.taskKey)

	pretty.Progress("Waiting for result...", func() {
		select {
		case result = <-listenResultsC:
		case err = <-resultsErrC:
		}
	})
	if err != nil {
		return err
	}
	if result.Error != "" {
		fmt.Printf("%s Task %s failed with an error:\n%s\n",
			pretty.FailSign,
			pretty.Fail(result.TaskKey),
			pretty.Fail(result.Error),
		)
	} else {
		fmt.Printf("%s Task %s returned output %s with data:\n%s\n",
			pretty.SuccessSign,
			pretty.Success(result.TaskKey),
			pretty.Colorize(pretty.FgCyan, result.OutputKey),
			pretty.ColorizeJSON(pretty.FgCyan, nil, false, []byte(result.OutputData)),
		)
	}
	return nil
}

func (c *serviceExecuteCmd) getTaskKey(s *coreapi.Service) error {
	keys := taskKeysFromService(s)

	if c.taskKey != "" {
		if !xstrings.SliceContains(keys, c.taskKey) {
			return fmt.Errorf("task %q does not exists on service", c.taskKey)
		}
		return nil
	}

	if len(keys) == 1 {
		c.taskKey = keys[0]
		return nil
	}

	if survey.AskOne(&survey.Select{
		Message: "Select the task to execute",
		Options: keys,
	}, &c.taskKey, nil) != nil {
		return errors.New("no task to execute")
	}
	return nil
}

func (c *serviceExecuteCmd) getData(s *coreapi.Service) (string, error) {
	if c.jsonFile != "" {
		return c.readFile()
	}

	// see if task has no inputs.
	noInput := false
	for _, task := range s.Tasks {
		if task.Key == c.taskKey {
			if len(task.Inputs) == 0 {
				noInput = true
			}
			break
		}
	}

	if noInput {
		if len(c.executeData) > 0 {
			return "", fmt.Errorf("task %q has no input but --data flag was supplied", c.taskKey)
		}
		return "{}", nil
	}

	if c.executeData != nil {
		castData, err := casting.TaskInputs(s, c.taskKey, c.executeData)
		if err != nil {
			return "", err
		}

		b, err := json.Marshal(castData)
		return string(b), err
	}

	if c.jsonFile == "" {
		if survey.AskOne(&survey.Input{Message: "Enter the filepath to the inputs"}, &c.jsonFile, nil) != nil {
			return "", errors.New("no filepath given")
		}
	}
	return c.readFile()
}

func (c *serviceExecuteCmd) readFile() (string, error) {
	content, err := xjson.ReadFile(c.jsonFile)
	return string(content), err
}

func taskKeysFromService(s *coreapi.Service) []string {
	var taskKeys []string
	for _, task := range s.Tasks {
		taskKeys = append(taskKeys, task.Key)
	}
	return taskKeys
}
