package routers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"gqlgen-lambda/database"
	"gqlgen-lambda/ent"
	"gqlgen-lambda/graph/resolvers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/", graphqlHandler(database.MysqlClient))
	r.GET("/query", playgroundHandler())
	return r
}

func graphqlHandler(client *ent.Client) gin.HandlerFunc {
	h := handler.NewDefaultServer(resolvers.NewSchema(client.Debug()))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
