package executor

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
)

type ErrExecute struct {
	SubError error
	Stdout   string
	Stderr   string
	ExitCode int
}

func (e ErrExecute) Error() string {
	return fmt.Sprintf("failed to execute script: '%v', stdout: '%s', stderr: '%s', exit code%d", e.SubError, e.Stdout, e.Stderr, e.ExitCode)
}

func Execute(script []byte) error {
	scriptFile := filepath.Join(os.TempDir(), randomString(16)+".sh")
	err := os.WriteFile(scriptFile, script, 0755)
	if err != nil {
		return fmt.Errorf("os.WriteFile: %w", err)
	}
	defer os.Remove(scriptFile)
	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)
	cmd := exec.Command(scriptFile)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err = cmd.Run()
	if err != nil {
		return &ErrExecute{
			SubError: err,
			Stdout:   stdout.String(),
			Stderr:   stderr.String(),
			ExitCode: cmd.ProcessState.ExitCode(),
		}
	}
	return nil
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

