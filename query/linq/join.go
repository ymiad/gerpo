package linq

import (
	"context"
	"strings"

	"github.com/insei/gerpo/sql"
)

func NewJoinBuilder(core *CoreBuilder) *JoinBuilder {
	return &JoinBuilder{
		core: core,
	}
}

type JoinBuilder struct {
	core *CoreBuilder
	opts []func(*sql.StringJoinBuilder)
}

func (q *JoinBuilder) Apply(b *sql.StringJoinBuilder) {
	for _, opt := range q.opts {
		opt(b)
	}
}

func (q *JoinBuilder) LeftJoin(leftJoinFn func(ctx context.Context) string) {
	q.opts = append(q.opts, func(builder *sql.StringJoinBuilder) {
		builder.JOIN(func(ctx context.Context) string {
			leftLoinStr := strings.TrimSpace(leftJoinFn(ctx))
			if leftLoinStr != "" {
				return "LEFT JOIN " + leftLoinStr
			}
			return ""
		})
	})
}
