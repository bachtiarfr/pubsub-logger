package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/bachtiarfr/pubsub-logger/v2/internal/entity"
	"github.com/bachtiarfr/pubsub-logger/v2/pkg/builderx"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	loggerData := entity.Logger{
		ID:         	uuid.New(),
		ServiceName: 	"ServiceName",
		Request:    	map[string]any{},
		RequestVendor:  map[string]any{},
		Response:   	map[string]any{},
		ResponseVendor: map[string]any{},
		URL:        	"dummy",
		UrlVendor:      "dummy",
		Method:    	 	"POST",
		StatusCode: 	200,
		Event:      	"nodeflux_liveness",
		RequestID:  	"state.RequestID",
		Additional: 	map[string]any{},
		Tag:        	"",
		Error:      	map[string]any{},
		CreatedAt:  	time.Now().Local(),
	}

	dataJson, _ := builderx.StructToMap(loggerData, "json")

	errPub := PublishReport(ctx, dataJson)
	if errPub != nil {
		fmt.Printf("publish error:", errPub)
	}
	fmt.Printf("success publish")

}

func PublishReport(ctx context.Context, data map[string]interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("error marshal :", err)
	}

	projectID := "bachtiar-development"
	credentialFilePath := "config/bachtiar-development-73ca13e5c16e.json"
	topicID := "dev-logger-topic"

	client, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		return fmt.Errorf("pubsub: NewClient: %v", err)
	}
	defer client.Close()

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(d),
	})
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("pubsub: result.Get: %v", err)
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
	return nil
}