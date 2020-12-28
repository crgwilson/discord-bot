package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/crgwilson/discord-bot/pkg/bot"
	"github.com/crgwilson/discord-bot/pkg/discord"
)

const (
	CliName        = "botctl"
	CliVersion     = "0.0.1"
	CliDescription = "Manage the discord bot"
	CliHelp        = `botctl: Manage the discord bot

Usage:
	botctl -flag <value>

Available Flags:
	-f       File path leading to the bot.yml config
`
)

var ErrMissingRequiredArgument = errors.New("Required argument was not provided")

func pingPong(args []string) (string, error) {
	return "pong", nil
}

func main() {
	configFilePtr := flag.String("f", "", "The file path of the yaml file containing the bot config")
	flag.Parse()

	if len(*configFilePtr) == 0 {
		fmt.Print(CliHelp)
	}

	configFileContent, err := ioutil.ReadFile(*configFilePtr)
	botConfig, err := bot.NewConfig(configFileContent)
	if err != nil {
		panic(err)
	}

	discord, err := discord.NewSession(botConfig.Discord.Token)
	if err != nil {
		panic(err)
	}

	pingPongCommand := bot.NewCommand("ping", "do some useless crap", "", pingPong)
	botCommands := []*bot.Command{pingPongCommand}
	botRouter, _ := bot.NewMessageRouter(botCommands)
	bot := bot.NewBot(discord, botRouter, botConfig)

	err = bot.Run()
	if err != nil {
		fmt.Println(err)
	}
}
