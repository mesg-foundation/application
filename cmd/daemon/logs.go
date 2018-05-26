package daemon

import (
	// "bytes"
	// "context"
	// "fmt"
	// "os"
	// "time"

	// godocker "github.com/fsouza/go-dockerclient"
	// "github.com/logrusorgru/aurora"
	// "github.com/mesg-foundation/core/daemon"
	// "github.com/mesg-foundation/core/docker"

	"bytes"
	"fmt"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/mesg-foundation/core/daemon"
	"github.com/spf13/cobra"
)

// Logs the daemon
var Logs = &cobra.Command{
	Use:               "logs",
	Short:             "Show the daemon's logs",
	Run:               logsHandler,
	DisableAutoGenTag: true,
}

func logsHandler(cmd *cobra.Command, args []string) {
	isRunning, err := daemon.IsRunning()
	if err != nil {
		fmt.Println(aurora.Red(err))
		return
	}
	if isRunning {
		var stream bytes.Buffer
		err = daemon.Logs(&stream)
		if err != nil {
			fmt.Println(aurora.Red(err))
			return
		}

		buf := make([]byte, 1024)
		for {
			n, _ := stream.Read(buf)
			if n != 0 {
				fmt.Print(string(buf[:n]))
			}
			time.Sleep(500 * time.Millisecond)
		}
	} else {
		fmt.Println(aurora.Brown("Daemon is stopped"))
	}
}
