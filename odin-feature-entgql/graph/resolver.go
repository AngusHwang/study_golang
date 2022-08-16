package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"odin/ent"
	"odin/graph/directives"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your mail, add any dependencies you require here.

type Resolver struct {
	client *ent.Client
}

func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{client: client},
		Directives: DirectiveRoot{
			Authorized: directives.Authorized,
		},
	})
}
