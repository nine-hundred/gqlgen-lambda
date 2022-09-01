package routers

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"gqlgen-lambda/database"
	"gqlgen-lambda/ent"
	"gqlgen-lambda/graph/resolvers"
	"log"
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

var ginLambda *ginadapter.GinLambda

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		log.Fatalf("gin cold start")
		r := gin.Default()
		r.POST("/", graphqlHandler(database.MysqlClient))
		r.GET("/query", playgroundHandler())
		ginLambda = ginadapter.New(r)
	}

	return ginLambda.ProxyWithContext(ctx, req)
}
