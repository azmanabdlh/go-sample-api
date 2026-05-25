package provider

import (
	"context"
	"errors"
	"sync"
)

type MemoryTokenProvider struct {
	mu     sync.RWMutex
	tokens map[string]bool
}

func NewMemoryTokenProvider() *MemoryTokenProvider {
	return &MemoryTokenProvider{
		tokens: map[string]bool{
			"bypass-token": true,
		},
	}
}

func (m *MemoryTokenProvider) ValidateToken(
	ctx context.Context,
	token string,
) error {
	defer m.mu.Unlock()
	m.mu.Lock()
	_, ok := m.tokens[token]
	if !ok {
		return errors.New("invalid token")
	}

	return nil
}
