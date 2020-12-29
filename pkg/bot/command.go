package bot

type commandCallback func(bot *Bot, args []string) (string, error)

type Command struct {
	Name        string
	Description string
	Usage       string
	Aliases     []string
	Callback    commandCallback
}

func (c Command) Call(b *Bot, args []string) (string, error) {
	val, err := c.Callback(b, args)
	return val, err
}

func NewCommand(name, description, usage string, aliases []string, callback commandCallback) *Command {
	command := Command{
		Name:        name,
		Description: description,
		Usage:       usage,
		Aliases:     aliases,
		Callback:    callback,
	}

	return &command
}
