// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package igvmmeasure

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

type MockCmd struct {
	output string
	err    error
}

func (m *MockCmd) CombinedOutput() ([]byte, error) {
	return []byte(m.output), m.err
}

func MockExecCommand(output string, err error) func(name string, arg ...string) *MockCmd {
	return func(name string, arg ...string) *MockCmd {
		return &MockCmd{
			output: output,
			err:    err,
		}
	}
}

func TestNewIgvmMeasurement(t *testing.T) {
	_, err := NewIgvmMeasurement("", nil, nil)
	if err == nil {
		t.Errorf("expected error for empty pathToFile, got nil")
	}

	igvm, err := NewIgvmMeasurement("/valid/path", nil, nil)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if igvm == nil {
		t.Errorf("expected non-nil IgvmMeasurement")
	}
}

func TestIgvmMeasurement_Run_Success(t *testing.T) {
	mockOutput := "measurement successful" // Ensure it's a **single-line output**

	m := &IgvmMeasurement{
		pathToFile: "/valid/path",
		execCommand: func(name string, arg ...string) *exec.Cmd {
			cmd := exec.Command("sh", "-c", "echo '"+mockOutput+"'") // Single line output
			return cmd
		},
	}

	err := m.Run("/mock/igvmBinary")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestIgvmMeasurement_Run_Error(t *testing.T) {
	mockOutput := "some error occurred"

	m := &IgvmMeasurement{
		pathToFile: "/invalid/path",
		execCommand: func(name string, arg ...string) *exec.Cmd {
			cmd := exec.Command("sh", "-c", "echo '"+mockOutput+"' && echo 'extra line' && exit 1") // Simulate multiline error
			return cmd
		},
	}

	err := m.Run("/mock/igvmBinary")

	if err == nil {
		t.Errorf("expected an error, got nil")
	} else if !strings.Contains(err.Error(), "error: "+mockOutput) {
		t.Errorf("expected error message to contain 'error: %s', got: %s", mockOutput, err)
	}
}

func TestIgvmMeasurement_Stop_NoProcess(t *testing.T) {
	m := &IgvmMeasurement{}

	err := m.Stop()
	if err == nil || err.Error() != "no running process to stop" {
		t.Errorf("expected 'no running process to stop' error, got: %v", err)
	}
}

func TestIgvmMeasurement_Stop_ProcessNil(t *testing.T) {
	m := &IgvmMeasurement{
		cmd: &exec.Cmd{},
	}

	err := m.Stop()
	if err == nil || err.Error() != "no running process to stop" {
		t.Errorf("expected 'no running process to stop' error, got: %v", err)
	}
}

func TestIgvmMeasurement_Stop_Success(t *testing.T) {
	process, err := os.StartProcess("/bin/sleep", []string{"sleep", "10"}, &os.ProcAttr{})
	if err != nil {
		t.Fatalf("failed to start mock process: %v", err)
	}
	defer func() {
		if err := process.Kill(); err != nil {
			t.Logf("Failed to kill process: %v", err)
		}
	}()

	m := &IgvmMeasurement{
		cmd: &exec.Cmd{Process: process},
	}

	err = m.Stop()
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}
