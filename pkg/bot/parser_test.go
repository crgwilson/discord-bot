package bot

import (
	"testing"
)

const (
	testPrefix      string = "!"
	testCommand     string = "testcommand"
	testCommandArgs string = " here are some args"
)

func TestParser(t *testing.T) {
	cases := []struct {
		Name             string
		InputMessage     string
		InputPrefix      string
		ExpectedCommand  string
		ExpectedArgCount int
	}{
		{
			"messages with no args should parse correctly",
			testPrefix + testCommand,
			testPrefix,
			testCommand,
			0,
		},
		{
			"messages with args should parse correctly",
			testPrefix + testCommand + testCommandArgs,
			testPrefix,
			testCommand,
			4,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			parsed, err := ParseMessage(test.InputMessage, test.InputPrefix)

			if err != nil {
				t.Errorf("got %v, want no error", err)
			}

			if parsed.RequestPath != test.ExpectedCommand {
				t.Errorf("got %q, want %q", parsed.RequestPath, test.ExpectedCommand)
			}

			argCount := len(parsed.RequestArgs)
			if argCount != test.ExpectedArgCount {
				t.Errorf("got %d args, expected 0", argCount)
			}
		})
	}

	t.Run("invalid message prefixes should return an error", func(t *testing.T) {
		_, err := ParseMessage(testCommand, "")

		if err != ErrInvalidCommandPrefix {
			t.Errorf("got %v, want %v", err, ErrInvalidCommandPrefix)
		}
	})

	t.Run("messages without prefixes should return an error", func(t *testing.T) {
		_, err := ParseMessage(testCommand, testPrefix)

		if err != ErrCommandPrefixNotFound {
			t.Errorf("got %v, want %v", err, ErrCommandPrefixNotFound)
		}
	})
}
