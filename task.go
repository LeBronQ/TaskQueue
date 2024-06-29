package tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/LeBronQ/kdtree"
	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	TypeKDtreeDelivery = "kdtree:deliver"
)

type KDtreeDeliveryPayload struct {
	Tree *kdtree.KDTree
}

//----------------------------------------------
// Write a function NewXXXTask to create a task.
// A task consists of a type and a payload.
//----------------------------------------------

func NewKDtreeDeliveryTask(t *kdtree.KDTree) (*asynq.Task, error) {
	payload, err := json.Marshal(KDtreeDeliveryPayload{Tree: t})
	if err != nil {
		return nil, err
	}
	fmt.Print(payload)
	return asynq.NewTask(TypeKDtreeDelivery, payload), nil
}

//---------------------------------------------------------------
// Write a function HandleXXXTask to handle the input task.
// Note that it satisfies the asynq.HandlerFunc interface.
//
// Handler doesn't need to be a function. You can define a type
// that satisfies asynq.Handler interface. See examples below.
//---------------------------------------------------------------

func HandleKDtreeDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p KDtreeDeliveryPayload
	fmt.Print("payload:", t.Payload())
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	fmt.Println("kdtree:", p.Tree)
	// Email delivery code ...
	return nil
}
