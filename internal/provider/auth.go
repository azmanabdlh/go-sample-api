package provider

import "context"

type Adapter interface {
	ValidateToken(
		ctx context.Context,
		token string,
	) error
}
