package bot

import (
	"errors"
	"strings"
)

var ErrInvalidCommandPrefix = errors.New("Command Prefix must be a non-empty string")
var ErrCommandPrefixNotFound = errors.New("The given command prefix is not present in the provided message")

type ParsedMessage struct {
	RequestPath string
	RequestArgs []string
}

func ParseMessage(message, messagePrefix string) (*ParsedMessage, error) {
	prefixLength := len(messagePrefix)

	if prefixLength == 0 {
		return nil, ErrInvalidCommandPrefix
	}

	if message[:prefixLength] != messagePrefix {
		return nil, ErrCommandPrefixNotFound
	}

	// Remove the prefix from the message
	strippedMessage := message[prefixLength:]

	splitMessage := strings.Split(strippedMessage, " ")

	parsed := ParsedMessage{
		RequestPath: splitMessage[0],
		RequestArgs: splitMessage[1:],
	}

	return &parsed, nil
}
