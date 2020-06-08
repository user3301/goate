package store

import (
	"context"

	"github.com/user3301/grpclab/pkg/types"
)

type UserStorer interface {
	CreateUser(ctx context.Context, userDetails types.UserDetails) error
	Verify(ctx context.Context, userDetails types.UserDetails) (bool, string)
	UpdateUser(ctx context.Context, details types.UserDetails) (bool, error)
}
