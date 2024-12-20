package cache

import (
	"context"

	"github.com/insei/gerpo/cache/types"
)

var (
	ErrNotFound           = types.ErrNotFound
	ErrWrongConfiguration = types.ErrWrongConfiguration
)

type Source = types.Source

type ModelBundle interface {
	Clean(ctx context.Context)
	Get(ctx context.Context, statement string, statementArgs ...any) (any, error)
	Set(ctx context.Context, cache any, statement string, statementArgs ...any)
}
