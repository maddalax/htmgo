package process

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)
import "golang.org/x/sys/windows"

func KillProcess(process *os.Process) error {
	if process == nil {
		return nil
	}
	Run(fmt.Sprintf("taskkill /F /T /PID %s", strconv.Itoa(process.Pid)))
	time.Sleep(time.Millisecond * 50)
	return nil
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
