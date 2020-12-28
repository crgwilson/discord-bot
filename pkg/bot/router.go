package bot

import (
	"errors"
)

var ErrRouteNotFound = errors.New("Could not find provided route")
var ErrRouteAlreadyExists = errors.New("Provided route already exists")

type MessageRouter struct {
	Routes map[string]*Command
	Prefix string
}

func (r *MessageRouter) Find(routeName string) (*Command, error) {
	val, ok := r.Routes[routeName]

	if !ok {
		return nil, ErrRouteNotFound
	}

	return val, nil
}

func (r *MessageRouter) RouteMessage(message string) (string, error) {
	// TODO: fix this please
	parsedMessage, _ := ParseMessage(message, r.Prefix)

	targetPath, err := r.Find(parsedMessage.RequestPath)
	if err != nil {
		return "", err
	}

	routeCallback := targetPath.Callback
	// TODO: Dont just swallow this error
	result, err := routeCallback(parsedMessage.RequestArgs)

	return result, err
}

func (r *MessageRouter) RegisterRoute(command *Command) error {
	_, err := r.Find(command.Name)
	if err != ErrRouteNotFound {
		return ErrRouteAlreadyExists
	}

	r.Routes[command.Name] = command

	return nil
}

func NewMessageRouter(commands []*Command) (*MessageRouter, error) {
	routeMap := make(map[string]*Command)

	router := MessageRouter{
		Routes: routeMap,
		Prefix: "!",
	}

	for _, c := range commands {
		err := router.RegisterRoute(c)
		if err != nil {
			return nil, err
		}
	}

	// TODO: actually do some error handling here please
	return &router, nil
}
