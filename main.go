package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/t1tc01/test-directive/graph"
	"github.com/t1tc01/test-directive/middleware"
	"github.com/t1tc01/test-directive/resolver"
)

const defaultPort = "3000"

func main() {

	//
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	c := graph.Config{Resolvers: &resolver.Resolver{}}
	c.Directives.HasRole = resolver.CoreDirective.HasRole

	app := fiber.New()
	server := handler.NewDefaultServer(graph.NewExecutableSchema(c))
	app.Get("/", adaptor.HTTPHandler(playground.Handler("GraphQL playground", "/query")))
	app.All("/query", middleware.AuthMiddleware(), adaptor.HTTPHandler(server))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(app.Listen(":" + port))
}
