package book

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(
	ctx context.Context,
	title string,
	authorId string,
) (*Book, error) {
	book := Book{
		ID:       uuid.NewString(),
		Title:    title,
		AuthorID: authorId,
	}

	err := s.repo.Create(ctx, book)

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (s *Service) Update(
	ctx context.Context,
	id string,
	title string,
	authorId string,
) (*Book, error) {
	book, err := s.repo.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if book == nil {
		return nil, fmt.Errorf("error: %w", ErrBookNotFound)
	}

	book.Title = title
	book.AuthorID = authorId

	err = s.repo.Update(ctx, *book)

	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *Service) Delete(
	ctx context.Context,
	id string,
) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) FindByID(
	ctx context.Context,
	id string,
) (*Book, error) {
	book, err := s.repo.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if book == nil {
		return nil, fmt.Errorf("error: %w", ErrBookNotFound)
	}

	return book, nil
}

func (s *Service) Search(
	ctx context.Context,
	query string,
	limit int,
	page int,
) ([]Book, int, error) {
	offset := (page - 1) * limit

	return s.repo.Search(
		ctx,
		query,
		limit,
		offset,
	)
}
