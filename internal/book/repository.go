package book

import "context"

type Repository interface {
	Create(ctx context.Context, book Book) error
	Update(ctx context.Context, book Book) error
	Delete(ctx context.Context, id string) error

	FindByID(ctx context.Context, id string) (*Book, error)

	Search(
		ctx context.Context,
		query string,
		limit int,
		offset int,
	) ([]Book, int, error)
}
