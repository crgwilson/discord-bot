package bot

import (
	"testing"
)

const (
	invalidConfig string = `---
id: 12345
secret: 54321
token: 09876
`
	validConfig string = `---
discord:
  id: 12345
  secret: 54321
  token: 09876
`
)

func TestConfig(t *testing.T) {
	assertSuccess := func(t *testing.T, got error) {
		if got != nil {
			t.Errorf("expected success, got error %v", got)
		}
	}

	assertString := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("valid config file should return valid config struct", func(t *testing.T) {
		validConfigBytes := []byte(validConfig)
		config, err := NewConfig(validConfigBytes)

		assertSuccess(t, err)
		assertString(t, config.Discord.ClientId, "12345")
		assertString(t, config.Discord.ClientSecret, "54321")
		assertString(t, config.Discord.Token, "09876")
	})

	t.Run("invalid config file but valid yaml file should an empty struct", func(t *testing.T) {
		invalidConfigBytes := []byte(invalidConfig)
		_, err := NewConfig(invalidConfigBytes)

		assertSuccess(t, err)
	})
}
