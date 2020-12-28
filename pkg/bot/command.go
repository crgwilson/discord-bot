package bot

// import (
// 	"github.com/bwmarrin/discordgo"
// )

// type commandCallback func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) (string, error)
type commandCallback func(args []string) (string, error)

type Command struct {
	Name        string
	Description string
	Usage       string
	// Aliases     []string
	Callback commandCallback
}

func (c Command) Call(args []string) (string, error) {
	val, err := c.Callback(args)
	return val, err
}

func NewCommand(name, description, usage string, callback commandCallback) *Command {
	command := Command{
		Name:        name,
		Description: description,
		Usage:       usage,
		// Aliases:     aliases,
		Callback: callback,
	}

	return &command
}
