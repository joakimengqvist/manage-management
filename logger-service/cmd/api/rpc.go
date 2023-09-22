package main

import (
	"context"
	"log"
	"log-service/data"
	"time"
)

type RPCServer struct{}

type RPCPayload struct {
	Action   string
	Name     string
	Data     string
	From     string
	To       string
	Subject  string
	Message  string
	Email    string
	Password string
}

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Action:    payload.Action,
		Name:      payload.Name,
		Data:      payload.Data,
		From:      payload.From,
		To:        payload.To,
		Subject:   payload.Subject,
		Message:   payload.Message,
		Email:     payload.Email,
		Password:  payload.Password,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	*resp = "Processed payload vid RPC: Name" + payload.Name + "Action: " + payload.Action
	return nil
}
