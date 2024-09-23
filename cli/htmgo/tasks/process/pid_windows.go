package process

import (
	"os"
	"os/exec"
)
import "golang.org/x/sys/windows"

func KillProcess(process *os.Process) error {
	return process.Kill()
}

func PrepareCommand(command *exec.Cmd) {

}

func PidExists(pid int32) bool {
	var handle, err = windows.OpenProcess(windows.PROCESS_QUERY_INFORMATION, false, uint32(pid))
	if err != nil {
		return false
	}
	defer windows.CloseHandle(handle)

	var exitCode uint32
	err = windows.GetExitCodeProcess(handle, &exitCode)
	if err != nil {
		return false
	}

	return exitCode == 259
}
