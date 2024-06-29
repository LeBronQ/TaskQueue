package tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/LeBronQ/kdtree"
	"github.com/hibiken/asynq"
)

type KDTree kdtree.KDTree

// A list of task types.
const (
	TypeEmailDelivery = "email:deliver"
)

type KDtreeDeliveryPayload struct {
	tree *KDTree
}

//----------------------------------------------
// Write a function NewXXXTask to create a task.
// A task consists of a type and a payload.
//----------------------------------------------

func NewKDtreeDeliveryTask(t *KDTree) (*asynq.Task, error) {
	payload, err := json.Marshal(KDtreeDeliveryPayload{tree: t})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, payload), nil
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
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	fmt.Println("kdtree:", p.tree)
	// Email delivery code ...
	return nil
}
