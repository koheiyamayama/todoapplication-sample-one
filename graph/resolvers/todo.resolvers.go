package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"todoapplication-sample-one/bff/graph/model"
)

func (r *queryResolver) GetTaskByID(ctx context.Context, id string) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListTasksByUserID(ctx context.Context, userID string) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListTasks(ctx context.Context) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}
