package process

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"
)
import "golang.org/x/sys/windows"

func KillProcess(process CmdWithFlags) error {
	if process.Cmd == nil || process.Cmd.Process == nil {
		return nil
	}
	err := exec.Command("taskkill", "/F", "/T", "/PID", strconv.Itoa(process.Cmd.Process.Pid)).Run()
	if err != nil {
		fmt.Println(err)
	}
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
