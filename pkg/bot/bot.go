package bot

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var ErrRouteNotFound = errors.New("Could not find provided route")
var ErrRouteAlreadyExists = errors.New("Provided route already exists")

type Bot struct {
	Session       *discordgo.Session
	Filter        *MessageFilter
	CommandPrefix string
	NamedRoutes   map[string]*Command
	Config        Config
}

func (b *Bot) FindRoute(routeName string) (*Command, error) {
	val, ok := b.NamedRoutes[routeName]

	if !ok {
		return nil, ErrRouteNotFound
	}

	return val, nil
}

func (b *Bot) RouteMessage(message string) (string, error) {
	parsedMessage, _ := ParseMessage(message, b.CommandPrefix)

	targetPath, err := b.FindRoute(parsedMessage.RequestPath)
	if err != nil {
		return "", err
	}

	result, err := targetPath.Call(b, parsedMessage.RequestArgs)

	return result, err
}

func (b *Bot) RegisterRoute(command *Command) error {
	routesToAdd := append(command.Aliases, command.Name)

	for _, r := range routesToAdd {
		_, err := b.FindRoute(r)
		if err != ErrRouteNotFound {
			return ErrRouteAlreadyExists
		}

		b.NamedRoutes[r] = command
	}

	return nil
}

func (b *Bot) HandleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := m.Content

	routeMessage := b.Filter.Filter(m.Author.ID, m.ChannelID, b.CommandPrefix, message)
	if !routeMessage {
		return
	}

	result, err := b.RouteMessage(message)
	if err != nil {
		return
	}

	s.ChannelMessageSend(m.ChannelID, result)
}

func (b *Bot) Run() error {
	err := b.Session.Open()
	if err != nil {
		return err
	}

	fmt.Println("Bot is running...")
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-signalChannel

	b.Session.Close()
	return nil
}

func NewBot(session *discordgo.Session, filter *MessageFilter, config *Config) *Bot {
	routeMap := make(map[string]*Command)

	bot := Bot{
		Session:       session,
		Filter:        filter,
		CommandPrefix: "!",
		NamedRoutes:   routeMap,
		Config:        *config,
	}

	session.AddHandler(bot.HandleMessage)

	return &bot
}
