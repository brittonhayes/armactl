package reforger

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
)

// StopProcess attempts to gracefully stop a process by PID
func StopProcess(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("process %d not found: %w", pid, err)
	}

	// Send termination signal (cross-platform)
	var signal os.Signal
	if runtime.GOOS == "windows" {
		// Windows: Kill is the only reliable option
		signal = os.Kill
	} else {
		// Unix: Try SIGTERM first (graceful shutdown)
		signal = syscall.SIGTERM
	}

	if err := process.Signal(signal); err != nil {
		return fmt.Errorf("failed to stop process %d: %w", pid, err)
	}

	return nil
}
