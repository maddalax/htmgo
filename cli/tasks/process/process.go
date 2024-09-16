package process

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

var commands = make([]*exec.Cmd, 0)

func AppendRunning(cmd *exec.Cmd) {
	commands = append(commands, cmd)
}

func KillAll() {

	tries := 0
	for {
		tries++
		allFinished := true
		for _, cmd := range commands {
			if cmd.Process == nil {
				allFinished = false

				if tries > 50 {
					args := strings.Join(cmd.Args, " ")
					log.Printf("process %v is not running after 50 tries, breaking.\n", args)
					allFinished = true
					break
				} else {
					time.Sleep(time.Millisecond * 50)
					continue
				}
			}
		}
		if allFinished {
			break
		}
	}

	for _, command := range commands {
		pid := command.Process.Pid
		err := syscall.Kill(-pid, syscall.SIGKILL)
		if err != nil {
			continue
		}
	}

	for {
		finished := true
		for _, c := range commands {
			if c.Process == nil {
				continue
			}
			exists, err := PidExists(int32(c.Process.Pid))
			if err != nil {
				finished = false
			}
			if exists {
				syscall.Kill(-c.Process.Pid, syscall.SIGKILL)
				finished = false
			}
		}

		if finished {
			break
		} else {
			fmt.Printf("waiting for all processes to exit\n")
			time.Sleep(time.Millisecond * 5)
		}
	}

	commands = make([]*exec.Cmd, 0)
}

func PidExists(pid int32) (bool, error) {
	if pid <= 0 {
		return false, fmt.Errorf("invalid pid %v", pid)
	}
	proc, err := os.FindProcess(int(pid))
	if err != nil {
		return false, err
	}
	err = proc.Signal(syscall.Signal(0))
	if err == nil {
		return true, nil
	}
	if err.Error() == "os: process already finished" {
		return false, nil
	}
	var errno syscall.Errno
	ok := errors.As(err, &errno)
	if !ok {
		return false, err
	}
	switch errno {
	case syscall.ESRCH:
		return false, nil
	case syscall.EPERM:
		return true, nil
	}
	return false, err
}

func RunOrExit(command string) {
	_ = Run(command, true)
}

func RunMany(commands []string, exitOnError bool) error {
	for _, command := range commands {
		err := Run(command, false)
		if err != nil {
			if exitOnError {
				os.Exit(1)
			}
			return err
		}
	}
	return nil
}

func Run(command string, exitOnError bool) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	AppendRunning(cmd)
	err := cmd.Run()
	if err != nil {
		if strings.Contains(err.Error(), "signal: killed") {
			return nil
		}
		if exitOnError {
			log.Println(fmt.Sprintf("error: %v", err))
			os.Exit(1)
		}
		return err
	}
	return nil
}
