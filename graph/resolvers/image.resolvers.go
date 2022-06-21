package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"todoapplication-sample-one/bff/graph/model"
)

func (r *queryResolver) GetImage(ctx context.Context, typeArg *model.ImageType, objectID string) (*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}
