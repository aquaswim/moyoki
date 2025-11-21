package port

import (
	"context"
)

type CrudRepository[T any] interface {
	Find(ctx context.Context) ([]T, error)
	FindByID(ctx context.Context, id int) (*T, error)
	Save(ctx context.Context, route *T) error
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, route *T) error
}
