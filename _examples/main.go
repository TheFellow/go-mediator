package main

import (
	"context"
	"errors"
	"log"

	"github.com/TheFellow/go-mediator/internal/must"
	"github.com/TheFellow/go-mediator/mediator"
)

func main() {
	m, err := mediator.New(
		mediator.WithBehaviourFunc(
			func(ctx context.Context, cmd mediator.Message, next mediator.Next) (interface{}, error) {
				log.Println("Pre Process - 1!")
				next(ctx)
				log.Println("Post Process - 1")

				return nil, nil
			}), mediator.WithBehaviourFunc(
			func(ctx context.Context, cmd mediator.Message, next mediator.Next) (interface{}, error) {
				log.Println("Pre Process!- 2")
				next(ctx)
				log.Println("Post Process - 2")

				return nil, nil
			}),
		mediator.WithHandler(&FakeCommand{}, NewFakeCommandHandler()))

	must.NotFail(err)

	cmd := &FakeCommand{
		Name: "Emre",
	}
	ctx := context.Background()

	m.Send(ctx, cmd)
}

type FakeCommand struct {
	Name string
}

func (*FakeCommand) Key() string { return "FakeCommand" }

type FakeCommandHandler struct{}

func NewFakeCommandHandler() FakeCommandHandler {
	return FakeCommandHandler{}
}

func (FakeCommandHandler) Handle(_ context.Context, command mediator.Message) (interface{}, error) {
	cmd := command.(*FakeCommand)
	if cmd.Name == "" {
		return nil, errors.New("Name is empty")
	}
	log.Println("handling fake cmd")
	return nil, nil
}
