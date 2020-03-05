package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	authentication "authentication/proto/authentication"
)

type Authentication struct{}

func (e *Authentication) Handle(ctx context.Context, msg *authentication.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *authentication.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
