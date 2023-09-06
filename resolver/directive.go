package resolver

import (
	"context"
	"fmt"
	"golang-boilerplate/internal/util"

	"github.com/99designs/gqlgen/graphql"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func validationResolver(v *validator.Validate, vt ut.Translator, l *zap.Logger) func(ctx context.Context, obj interface{}, next graphql.Resolver, constrains string) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, constrains string) (interface{}, error) {
		val, err := next(ctx)
		if err != nil {
			l.Error("Getting error when extract values from context", zap.Error(err))
			return nil, util.WrapGQLError(ctx, err.Error(), "extract_values_from_context_has_failed")
		}

		fieldName := *graphql.GetPathContext(ctx).Field

		err = v.Var(val, constrains)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			return nil, fmt.Errorf("%s:%s", fieldName, validationErrors[0].Translate(vt))
		}

		return val, nil
	}
}
