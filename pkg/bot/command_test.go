package bot

import (
	"errors"
	"strings"
	"testing"
)

const (
	testName        string = "test"
	testAlias       string = "testAlias"
	testDescription string = "this is a unit test"
	testUsage       string = "test <SOME ARG>"
)

var testErr = errors.New("this is a test")

func joinSlice(b *Bot, args []string) (string, error) {
	val := strings.Join(args, ",")
	return val, nil
}

func returnAnError(args []string) (string, error) {
	return "", testErr
}

func TestCommand(t *testing.T) {
	testAliases := []string{testAlias}

	t.Run("successful command invocation", func(t *testing.T) {
		testCommand := NewCommand(testName, testDescription, testUsage, testAliases, joinSlice)

		if testCommand.Name != testName {
			t.Errorf("got %q, want %q", testCommand.Name, testName)
		}
		if testCommand.Description != testDescription {
			t.Errorf("got %q, want %q", testCommand.Description, testDescription)
		}
		if testCommand.Usage != testUsage {
			t.Errorf("got %q, want %q", testCommand.Usage, testDescription)
		}
	})
}
