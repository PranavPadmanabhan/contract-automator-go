package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AutoTask struct {
	ID         primitive.ObjectID `json:"_id"  bson:"_id,omitempty"`
	Address    string             `json:"address" bson:"address,omitempty"`
	Abi        string             `json:"abi" bson:"abi,omitempty"`
	Executions []string           `json:"executions" bson:"executions"`
	// ExecutorAddress string             `json:"executoraddress" bson:"executoraddress"`
	// ExecutorKey     string             `json:"executorkey" bson:"executorkey"`
}
