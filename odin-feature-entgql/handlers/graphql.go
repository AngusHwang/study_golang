package handlers

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"odin/ent"
	"odin/graph"
)

func GraphQL(client *ent.Client) gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewSchema(client))

	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "Authorization", c.GetHeader("Authorization"))
		h.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	}
}
