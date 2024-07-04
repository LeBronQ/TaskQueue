package tasks

import (
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	TypeKDtreeDelivery = "kdtree:deliver"
)

type TreeNodeData struct {
	ID int64
}

type DeliveryPoint struct {
	Coordinates []float64 `json:"coordinates"`
	ID          int64     `json:"id"`
}

type KDtreeDeliveryPayload struct {
	TreeNodes []DeliveryPoint
}

//----------------------------------------------
// Write a function NewXXXTask to create a task.
// A task consists of a type and a payload.
//----------------------------------------------

func NewKDtreeDeliveryTask(t []DeliveryPoint) (*asynq.Task, error) {
	payload, err := json.Marshal(KDtreeDeliveryPayload{TreeNodes: t})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeKDtreeDelivery, payload), nil
}
