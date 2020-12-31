package main

import (
	"strings"

	"github.com/crgwilson/discord-bot/pkg/bot"
)

func help(b *bot.Bot, args []string) (string, error) {
	messageBorder := "```"

	output := "Hi, I am a stupid bot that does stupid things...\n\n"
	output += messageBorder + "\nAvailable Bot Commands:\n\n"

	for k, v := range b.NamedRoutes {
		if k != v.Name {
			// This is an alias, so we don't want to print it out to avoid repeats
			continue
		}

		line := b.CommandPrefix + v.Name + "\t\t" + v.Description + "\n"

		if len(v.Aliases) > 0 {
			line += "\t!" + strings.Join(v.Aliases, "\n\t!")
		}

		line += "\n"
		output += line
	}
	output += messageBorder

	return output, nil
}

func NewHelpCommand() *bot.Command {
	command := bot.NewCommand(
		"help",
		"Describe all available bot commands",
		"",
		[]string{"h", "hlep"},
		help,
	)
	return command
}
