package util

import (
	"context"

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
