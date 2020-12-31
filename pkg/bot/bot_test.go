package bot

import (
	"strings"
	"testing"

	"github.com/bwmarrin/discordgo"
)

const (
	testCommandPrefix      string = "!"
	testCommandName        string = "test"
	testCommandDescription string = "test description"
	testCommandUsage       string = "test usage"
	testCommandAlias       string = "alias"
)

func callback(b *Bot, args []string) (string, error) {
	joinedSlice := strings.Join(args, ",")
	return joinedSlice, nil
}

func TestBot(t *testing.T) {
	testCommand := NewCommand(
		testCommandName,
		testCommandDescription,
		testCommandUsage,
		[]string{testCommandAlias},
		callback,
	)

	testSession := discordgo.Session{}

	filter := NewMessageFilter(
		make([]string, 0),
		make([]string, 0),
	)

	config := Config{}

	testBot := NewBot(&testSession, filter, &config)

	t.Run("adding new routes", func(t *testing.T) {
		err := testBot.RegisterRoute(testCommand)
		if err != nil {
			t.Errorf("got %v, expected no error", err)
		}
	})

	t.Run("find existing routes", func(t *testing.T) {
		result, err := testBot.FindRoute(testCommandName)

		if err != nil {
			t.Errorf("got %v, want no error", err)
		}

		if result != testCommand {
			t.Errorf("got %v, want %v", result, testCommand)
		}
	})

	t.Run("error out when the route does not exist", func(t *testing.T) {
		_, err := testBot.FindRoute("notreal")

		if err != ErrRouteNotFound {
			t.Errorf("got %v, want %v", err, ErrRouteNotFound)
		}
	})

	t.Run("route message to callback with no args", func(t *testing.T) {
		testMessage := testCommandPrefix + testCommandName

		result, err := testBot.RouteMessage(testMessage)
		if err != nil {
			t.Errorf("got %v, want no error", err)
		}

		if result != "" {
			t.Errorf("got %s, want empty string", result)
		}
	})

	t.Run("route message to callback with some args", func(t *testing.T) {
		testMessage := testCommandPrefix + testCommandName + " here are some args"
		expectedResult := "here,are,some,args"

		result, err := testBot.RouteMessage(testMessage)
		if err != nil {
			t.Errorf("got %v, want no error", err)
		}

		if result != expectedResult {
			t.Errorf("got %q, want %q", result, expectedResult)
		}
	})

	t.Run("route aliases should work too", func(t *testing.T) {
		testMessage := testCommandPrefix + testCommandAlias

		result, err := testBot.RouteMessage(testMessage)
		if err != nil {
			t.Errorf("got %v, want no error", err)
		}

		if result != "" {
			t.Errorf("got %s, want empty string", result)
		}
	})
}
