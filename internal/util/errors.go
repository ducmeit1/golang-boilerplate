package util

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func WrapGQLError(ctx context.Context, message string, code string) *gqlerror.Error {
	e := &gqlerror.Error{
		Message: message,
		Extensions: map[string]interface{}{
			"code": code,
		},
	}

	if ctx != nil {
		e.Path = graphql.GetPath(ctx)
	}

	return e
}

func WrapGQLInternalError(ctx context.Context) *gqlerror.Error {
	return WrapGQLError(ctx, "Internal Server Error", "INTERNAL_SERVER_ERROR")
}

func WrapGQLBadRequestError(ctx context.Context, format string, args ...interface{}) *gqlerror.Error {
	return WrapGQLError(ctx, fmt.Sprintf(format, args...), "BAD_REQUEST")
}

func WrapGQLUnauthorizedError(ctx context.Context) *gqlerror.Error {
	return WrapGQLError(ctx, "Unauthorized Request", "UNAUTHORIZED")
}

func WrapGQLNotFoundError(ctx context.Context) *gqlerror.Error {
	return WrapGQLError(ctx, "Not Found", "NOT_FOUND")
}
