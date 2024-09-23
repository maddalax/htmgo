package process

import (
	"log/slog"
	"os"
	"os/exec"
)

func RunOrExit(command string) {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		slog.Error("Error running command: ", slog.String("command", command))
		os.Exit(1)
	}
}
