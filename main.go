package bfr

import (
	"context"
	"encoding/json"

	"github.com/bachtiarfr/pubsub-logger/internal/entity"
	"github.com/bachtiarfr/pubsub-logger/pkg/googlepublisher"

	"github.com/pkg/errors"
)

func PublishReport(ctx context.Context, data entity.Logger) error {
	var r *googlepublisher.Client
	if r != nil {
		d, err := json.Marshal(data)
		if err != nil {
			return errors.Wrap(err, "data logger marshal failed")
		}

		return r.Publish(ctx, d)
	}
	return nil
}