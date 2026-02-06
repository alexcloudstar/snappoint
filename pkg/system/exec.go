package system

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

// CommandExecutor defines the interface for executing commands
type CommandExecutor interface {
	Execute(ctx context.Context, name string, args ...string) (string, error)
	IsAvailable(ctx context.Context, command string) bool
}

// RealExecutor implements CommandExecutor using actual system commands
type RealExecutor struct {
	Timeout time.Duration
}

// NewExecutor creates a new command executor with a default timeout
func NewExecutor() *RealExecutor {
	return &RealExecutor{
		Timeout: 30 * time.Second,
	}
}

// Execute runs a command with the given arguments and returns the output
func (e *RealExecutor) Execute(ctx context.Context, name string, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, e.Timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		if stderr.Len() > 0 {
			return "", fmt.Errorf("command '%s %v' failed: %w\nstderr: %s", name, args, err, stderr.String())
		}
		return "", fmt.Errorf("command '%s %v' failed: %w", name, args, err)
	}

	return stdout.String(), nil
}

// IsAvailable checks if a command is available on the system
func (e *RealExecutor) IsAvailable(ctx context.Context, command string) bool {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := exec.LookPath(command)
	return err == nil
}
