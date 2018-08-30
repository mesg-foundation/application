package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/mesg-foundation/core/cmd/utils"
	"github.com/mesg-foundation/core/interface/grpc/core"
	"github.com/mesg-foundation/core/x/xsignal"
	"github.com/spf13/cobra"
)

// Dev command will run the service from a path in dev mode
// It will also listen for all events and outputs from the tasks
var Dev = &cobra.Command{
	Use:               "dev",
	Short:             "Run a service in development mode",
	Example:           "mesg-core service dev PATH",
	Run:               devHandler,
	DisableAutoGenTag: true,
}

func init() {
	Dev.Flags().StringP("event-filter", "e", "*", "Only log the data of the given event")
	Dev.Flags().StringP("task-filter", "t", "", "Only log the result of the given task")
	Dev.Flags().StringP("output-filter", "o", "", "Only log the data of the given output of a task result. If set, you also need to set the task in --task-filter")
}

func devHandler(cmd *cobra.Command, args []string) {
	serviceID, isValid, err := createService(defaultPath(args))
	utils.HandleError(err)
	if !isValid {
		os.Exit(1)
	}
	fmt.Printf("%s Service started with success\n", aurora.Green("✔"))
	fmt.Printf("Service ID: %s\n", aurora.Bold(serviceID))

	go listenEvents(serviceID, cmd.Flag("event-filter").Value.String())
	go listenResults(serviceID, cmd.Flag("task-filter").Value.String(), cmd.Flag("output-filter").Value.String())

	closeReaders := showLogs(serviceID, "*")
	defer closeReaders()

	<-xsignal.WaitForInterrupt()

	utils.ShowSpinnerForFunc(utils.SpinnerOptions{Text: "Stopping service..."}, func() {
		cli().DeleteService(context.Background(), &core.DeleteServiceRequest{ // Delete service. This will automatically stop the service too
			ServiceID: serviceID,
		})
	})
}

func createService(path string) (serviceID string, isValid bool, err error) {
	serviceID, isValid, err = deployService(path)
	if !isValid || err != nil {
		return "", isValid, err
	}

	utils.ShowSpinnerForFunc(utils.SpinnerOptions{Text: "Starting service..."}, func() {
		_, err = cli().StartService(context.Background(), &core.StartServiceRequest{
			ServiceID: serviceID,
		})
	})
	return serviceID, true, err
}

func listenEvents(serviceID string, filter string) {
	stream, err := cli().ListenEvent(context.Background(), &core.ListenEventRequest{
		ServiceID:   serviceID,
		EventFilter: filter,
	})
	utils.HandleError(err)
	fmt.Println(aurora.Cyan("Listening for events from the service..."))
	for {
		event, err := stream.Recv()
		if err != nil {
			log.Println(aurora.Red(err))
			return
		}
		log.Println("Receive event", aurora.Green(event.EventKey), "with data", formatJSON(event.EventData))
	}
}

func listenResults(serviceID string, result string, output string) {
	stream, err := cli().ListenResult(context.Background(), &core.ListenResultRequest{
		ServiceID:    serviceID,
		TaskFilter:   result,
		OutputFilter: output,
	})
	utils.HandleError(err)
	fmt.Println(aurora.Cyan("Listening for results from the service..."))
	for {
		result, err := stream.Recv()
		if err != nil {
			log.Println(aurora.Red(err))
			return
		}
		log.Println("Receive result", aurora.Green(result.TaskKey), aurora.Cyan(result.OutputKey), "with data", formatJSON(result.OutputData))
	}
}

func formatJSON(data string) string {
	var decoded map[string]interface{}
	if err := json.Unmarshal([]byte(data), &decoded); err != nil {
		return data
	}
	var outputs []string
	for key, value := range decoded {
		outputs = append(outputs, fmt.Sprintf("%v = %v", aurora.Cyan(key), value))
	}
	return strings.Join(outputs, ", ")
}
