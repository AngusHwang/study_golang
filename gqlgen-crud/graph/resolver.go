package graph

import (
	"gqlgen-crud/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos   []*model.Todo
	members []*model.Member
}
