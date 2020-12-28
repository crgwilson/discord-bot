package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type messageHandler func(s *discordgo.Session, m *discordgo.MessageCreate)

type Bot struct {
	Session *discordgo.Session
	Router  *MessageRouter
	Config  Config
}

func (b *Bot) HandleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore my own messages
	if m.Author.ID == s.State.User.ID {
		return
	}

	message := m.Content

	// Commands start with the specified prefix
	prefixLength := len(b.Router.Prefix)
	if message[:prefixLength] != b.Router.Prefix {
		return
	}

	result, err := b.Router.RouteMessage(message)
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

func NewBot(session *discordgo.Session, router *MessageRouter, config *Config) *Bot {
	bot := Bot{
		Session: session,
		Router:  router,
		Config:  *config,
	}

	session.AddHandler(bot.HandleMessage)

	return &bot
}
