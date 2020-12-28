package bot

import (
	"errors"
	"strings"
	"testing"
)

const (
	testName        string = "test"
	testDescription string = "this is a unit test"
	testUsage       string = "test <SOME ARG>"
)

var testErr = errors.New("this is a test")

func joinSlice(args []string) (string, error) {
	val := strings.Join(args, ",")
	return val, nil
}

func returnAnError(args []string) (string, error) {
	return "", testErr
}

func TestCommand(t *testing.T) {
	t.Run("successful command invocation", func(t *testing.T) {
		testCommand := NewCommand(testName, testDescription, testUsage, joinSlice)

		if testCommand.Name != testName {
			t.Errorf("got %q, want %q", testCommand.Name, testName)
		}
		if testCommand.Description != testDescription {
			t.Errorf("got %q, want %q", testCommand.Description, testDescription)
		}
		if testCommand.Usage != testUsage {
			t.Errorf("got %q, want %q", testCommand.Usage, testDescription)
		}

		testArgs := []string{"this", "is", "a", "unit", "test"}
		output, err := testCommand.Call(testArgs)

		if err != nil {
			t.Errorf("got %v, want no error", err)
		}

		testOutput := strings.Join(testArgs, ",")
		if output != testOutput {
			t.Errorf("got %q, want %q", output, testOutput)
		}
	})

	t.Run("failing command invocation", func(t *testing.T) {
		testCommand := NewCommand(testName, testDescription, testUsage, returnAnError)
		args := make([]string, 0)
		_, err := testCommand.Call(args)

		if err != testErr {
			t.Errorf("got %v, want %v", err, testErr)
		}
	})
}
