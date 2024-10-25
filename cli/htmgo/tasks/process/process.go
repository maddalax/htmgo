package process

import (
	"fmt"
	"github.com/maddalax/htmgo/cli/htmgo/internal"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"
)

type CmdWithFlags struct {
	Flags []RunFlag
	Name  string
	Cmd   *exec.Cmd
}

type RawCommand struct {
	Name  string
	Args  string
	Flags []RunFlag
}

func NewRawCommand(name string, args string, flags ...RunFlag) RawCommand {
	if name == "" {
		name = args
	}
	c := RawCommand{Name: name, Args: args, Flags: flags}
	if c.Flags == nil {
		c.Flags = make([]RunFlag, 0)
	}
	return c
}

var workingDir string
var commands = make(map[string]CmdWithFlags)

func AppendRunning(cmd *exec.Cmd, raw RawCommand) {
	slog.Debug("running", slog.String("command", strings.Join(cmd.Args, " ")),
		slog.String("dir", cmd.Dir),
		slog.String("cwd", GetWorkingDir()))

	commands[raw.Name] = CmdWithFlags{Flags: raw.Flags, Name: raw.Name, Cmd: cmd}
}

func GetWorkingDir() string {
	if workingDir == "" {
		wd, _ := os.Getwd()
		return wd
	}
	return workingDir
}

func SetWorkingDir(dir string) {
	workingDir = dir
}

func GetPathRelativeToCwd(path string) string {
	return filepath.Join(GetWorkingDir(), path)
}

func shouldSkipKilling(flags []RunFlag, skipFlag []RunFlag) bool {
	for _, flag := range flags {
		if slices.Contains(skipFlag, flag) {
			return true
		}
	}
	return false
}

func StartLogger() {
	if internal.GetLogLevel() != slog.LevelDebug {
		return
	}
	go func() {
		for {
			time.Sleep(time.Second * 5)
			items := make([]map[string]string, 0)
			for _, cmd := range commands {
				data := make(map[string]string)
				data["command"] = fmt.Sprintf("%s %s", cmd.Cmd.Path, strings.Join(cmd.Cmd.Args, " "))
				if cmd.Cmd.Process != nil {
					data["pid"] = fmt.Sprintf("%d", cmd.Cmd.Process.Pid)
				}
				items = append(items, data)
			}

			fmt.Printf("Running processes:\n")
			for i, item := range items {
				fmt.Printf("%d: %+v\n", i, item)
			}
			fmt.Printf("\n")
		}
	}()
}

func GetProcessByName(name string) *CmdWithFlags {
	for _, cmd := range commands {
		if cmd.Name == name {
			return &cmd
		}
	}
	return nil
}

func OnShutdown() {
	// request for shutdown
	for _, cmd := range commands {
		if cmd.Cmd != nil && cmd.Cmd.Process != nil {
			cmd.Cmd.Process.Signal(os.Interrupt)
		}
	}
	// give it a second
	time.Sleep(time.Second * 1)
	// force kill
	KillAll()
}

func KillAll(skipFlag ...RunFlag) {

	tries := 0
	updatedCommands := make(map[string]CmdWithFlags)
	for {
		tries++
		allFinished := true
		for _, cmd := range commands {
			if cmd.Cmd.Process == nil {
				allFinished = false
				if tries > 50 {
					args := strings.Join(cmd.Cmd.Args, " ")
					slog.Debug("process is not running after 50 tries, breaking.", slog.String("command", args))
					allFinished = true
					break
				} else {
					time.Sleep(time.Millisecond * 50)
					continue
				}
			} else {
				updatedCommands[cmd.Name] = cmd
			}
		}
		if allFinished {
			break
		}
	}

	commands = make(map[string]CmdWithFlags)
	for _, command := range updatedCommands {
		if command.Cmd != nil && command.Cmd.Process != nil {
			commands[command.Name] = command
		}
	}

	for _, command := range commands {
		if shouldSkipKilling(command.Flags, skipFlag) {
			continue
		}
		err := KillProcess(command)
		if err != nil {
			continue
		}
	}

	for {
		finished := true
		for _, c := range commands {
			if c.Cmd.Process == nil {
				continue
			}
			if shouldSkipKilling(c.Flags, skipFlag) {
				continue
			}
			exists := PidExists(int32(c.Cmd.Process.Pid))
			if exists {
				KillProcess(c)
				finished = false
			}
		}

		if finished {
			break
		} else {
			slog.Debug("waiting for all processes to exit\n")
			time.Sleep(time.Millisecond * 5)
		}
	}

	commands = make(map[string]CmdWithFlags)
	slog.Debug("all processes killed\n")
}

func RunOrExit(command RawCommand) {
	command.Flags = append(command.Flags, ExitOnError)
	_ = Run(command)
}

type RunFlag int

const (
	ExitOnError RunFlag = iota
	Silent
	KillOnlyOnExit
)

func RunMany(commands ...RawCommand) error {
	for _, command := range commands {
		err := Run(command)
		if err != nil {
			if slices.Contains(command.Flags, ExitOnError) {
				os.Exit(1)
			}
			return err
		}
	}
	return nil
}

var mutex = &sync.Mutex{}

func Run(command RawCommand) error {
	mutex.Lock()

	parts := strings.Fields(command.Args)

	args := make([]string, 0)
	if len(parts) > 1 {
		args = parts[1:]
	}

	path := parts[0]

	existing := GetProcessByName(command.Name)

	if existing != nil {
		slog.Debug("process already running, killing it", slog.String("command", command.Name))
		KillProcess(*existing)
		time.Sleep(time.Millisecond * 50)
	} else {
		slog.Debug("no existing process found for %s, safe to run...", slog.String("command", command.Name))
	}

	cmd := exec.Command(path, args...)

	PrepareCommand(cmd)

	if slices.Contains(command.Flags, Silent) {
		cmd.Stdout = nil
		cmd.Stderr = nil
	} else {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if workingDir != "" {
		cmd.Dir = workingDir
	}

	AppendRunning(cmd, command)

	mutex.Unlock()

	err := cmd.Run()

	slog.Debug("command finished",
		slog.String("command", command.Name),
		slog.String("args", command.Args),
		slog.String("dir", cmd.Dir),
		slog.String("cwd", GetWorkingDir()),
		slog.String("error", fmt.Sprintf("%v", err)))

	delete(commands, command.Name)

	if err == nil {
		return nil
	}

	if strings.Contains(err.Error(), "signal: killed") {
		return nil
	}

	if slices.Contains(command.Flags, ExitOnError) {
		slog.Error("Error running command: ",
			slog.String("error", err.Error()),
			slog.String("command", command.Name))
		os.Exit(1)
	}

	return err
}
