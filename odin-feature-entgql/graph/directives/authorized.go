package directives

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"odin/securities"
	"strings"
)

var UserCtxKey = "user"

func Authorized(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	authorization, ok := ctx.Value("Authorization").(string)
	if !ok {
		return nil, errors.New("permission denied")
	}

	tokens := strings.Split(authorization, " ")
	if tokens[0] != "Bearer" || len(tokens) <= 1 {
		return nil, errors.New("permission denied")
	}

	claim, err := securities.Parse(tokens[1])
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, UserCtxKey, claim)

	return next(ctx)
}
