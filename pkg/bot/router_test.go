package bot

import (
	"testing"
)

const (
	testMessage string = "!test"
	testResult  string = "I made it"
)

func doNothing(args []string) (string, error) {
	return testResult, nil
}

func TestMessageRouter(t *testing.T) {
	command := NewCommand("test", "this is a test command", "test usage", doNothing)
	commands := []*Command{command}

	t.Run("message router with command", func(t *testing.T) {
		testRouter, err := NewMessageRouter(commands)
		if err != nil {
			t.Errorf("got %v, expected no error", err)
		}

		findResult, err := testRouter.Find(command.Name)
		if err != nil {
			t.Errorf("got %v, expected no error", err)
		}
		if findResult != command {
			t.Errorf("got %v, want %v", findResult, command)
		}

		_, err = testRouter.Find("doesntexist")
		if err != ErrRouteNotFound {
			t.Errorf("got %v, want %v", err, ErrRouteNotFound)
		}

		err = testRouter.RegisterRoute(command)
		if err != ErrRouteAlreadyExists {
			t.Errorf("got %v, want %v", err, ErrRouteAlreadyExists)
		}

		routed, err := testRouter.RouteMessage(testMessage)
		if routed != testResult {
			t.Errorf("got %q, want %q", routed, testResult)
		}
	})
}
