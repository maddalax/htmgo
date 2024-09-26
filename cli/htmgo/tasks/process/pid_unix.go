//go:build linux || darwin

package process

import (
	"errors"
	"log/slog"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func KillProcess(process CmdWithFlags) error {
	if process.Cmd == nil || process.Cmd.Process == nil {
		return nil
	}
	slog.Debug("killing process",
		slog.String("name", process.Name),
		slog.Int("pid", process.Cmd.Process.Pid))
	_ = syscall.Kill(-process.Cmd.Process.Pid, syscall.SIGKILL)
	time.Sleep(time.Millisecond * 50)
	return nil
}

func PrepareCommand(command *exec.Cmd) {
	command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}

func PidExists(pid int32) bool {
	if pid <= 0 {
		return false
	}
	proc, err := os.FindProcess(int(pid))
	if err != nil {
		return false
	}
	err = proc.Signal(syscall.Signal(0))
	if err == nil {
		return true
	}
	if err.Error() == "os: process already finished" {
		return false
	}
	var errno syscall.Errno
	ok := errors.As(err, &errno)
	if !ok {
		return false
	}
	switch errno {
	case syscall.ESRCH:
		return false
	case syscall.EPERM:
		return true
	}
	return false
}
