package module

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetDependencyPath(dep string) string {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}", dep)
	// Run the command and capture the output
	output, err := cmd.CombinedOutput() // Use CombinedOutput to capture both stdout and stderr
	if err != nil {
		fmt.Printf("Command execution failed: %v\n", err)
	}

	// Convert output to string
	dir := strings.TrimSuffix(string(output), "\n")
	if strings.Contains(dir, "not a known dependency") {
		return dep
	}
	return dir
}
