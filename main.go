package bfr

import (
	"context"

	"github.com/bachtiarfr/pubsub-logger/v2/internal/entity"
	"github.com/bachtiarfr/pubsub-logger/v2/pkg/googlepublisher"
)

type PublisherService interface {
	PublishReport(ctx context.Context, data entity.Logger) error
}

type LoggerService struct {
	publisher *googlepublisher.Client
}

func NewLoggerService(publisher *googlepublisher.Client) *LoggerService {
	return &LoggerService{publisher: publisher}
}

func PublishReport(ctx context.Context, data []byte) error {
	var r *googlepublisher.Client
	if r != nil {
		return r.Publish(ctx, data)
	}
	return nil
}