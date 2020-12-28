package discord

import (
	"github.com/bwmarrin/discordgo"
)

func NewSession(token string) (*discordgo.Session, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	// Only receive messages for now
	session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	return session, nil
}
