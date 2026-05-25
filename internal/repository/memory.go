package repository

import (
	"context"
	"strings"
	"sync"

	"github.com/azmanabdlh/go-sample-api/internal/book"
)

type MemoryStore struct {
	mu    sync.RWMutex
	books map[string]book.Book
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		books: map[string]book.Book{},
	}
}

func (m *MemoryStore) Create(
	ctx context.Context,
	book book.Book,
) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.books[book.ID] = book

	return nil
}

func (m *MemoryStore) Update(
	ctx context.Context,
	book book.Book,
) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.books[book.ID]

	if !ok {
		return nil
	}

	m.books[book.ID] = book

	return nil
}

func (m *MemoryStore) Delete(
	ctx context.Context,
	id string,
) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.books, id)

	return nil
}

func (m *MemoryStore) FindByID(
	ctx context.Context,
	id string,
) (*book.Book, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	book, ok := m.books[id]

	if !ok {
		return nil, nil
	}

	return &book, nil
}

func (m *MemoryStore) Search(
	ctx context.Context,
	query string,
	limit int,
	offset int,
) ([]book.Book, int, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var result []book.Book

	for _, item := range m.books {
		if query != "" {
			if !strings.Contains(
				strings.ToLower(item.Title),
				strings.ToLower(query),
			) {
				continue
			}
		}

		result = append(result, item)
	}

	total := len(result)

	start := offset
	end := offset + limit

	if start > total {
		return []book.Book{}, total, nil
	}

	if end > total {
		end = total
	}

	return result[start:end], total, nil
}
