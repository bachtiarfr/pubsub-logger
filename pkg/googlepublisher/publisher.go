package googlepublisher

import (
	"context"

	"github.com/pkg/errors"

	"google.golang.org/api/option"

	"cloud.google.com/go/pubsub"
)

type Config struct {
	ProjectID      string
	TopicID        string
	CredentialFile string
	Enable         bool
}

type Client struct {
	client *pubsub.Client
	topic  *pubsub.Topic
	Enable bool
}

func New(ctx context.Context, cfg Config) (*Client, error) {
	client, err := pubsub.NewClient(ctx, "bachtiar-development", option.WithCredentialsFile("config/bachtiar-development-73ca13e5c16e.json"))
	if err != nil {
		return nil, errors.Wrap(err, "new client pubsub error")
	}

	topic := client.Topic("dev-logger-topic")

	return &Client{
		client: client,
		topic:  topic,
		Enable: true,
	}, nil
}

type PublishOptions struct {
	Attributes map[string]string
}

type PublishOpts func(*PublishOptions)

func WithAttributes(attributes map[string]string) PublishOpts {
	return func(options *PublishOptions) {
		options.Attributes = attributes
	}
}

func (c *Client) Publish(ctx context.Context, message []byte, opts ...PublishOpts) error {
	publishOpts := PublishOptions{Attributes: make(map[string]string)}
	for _, opt := range opts {
		opt(&publishOpts)
	}

	return c.publish(ctx, &pubsub.Message{
		Data:       message,
		Attributes: publishOpts.Attributes,
	})
}

func (c *Client) publish(ctx context.Context, message *pubsub.Message) error {
	if c.Enable {
		result := c.topic.Publish(ctx, message)

		_, err := result.Get(ctx)
		if err != nil {
			return errors.Wrap(err, "error publish message")
		}
	}

	return nil
}
