package resolver

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/t1tc01/test-directive/common/jwt"
	"github.com/t1tc01/test-directive/graph"
	"github.com/t1tc01/test-directive/middleware"
	"github.com/t1tc01/test-directive/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func NewSchema() graphql.ExecutableSchema {
	return graph.NewExecutableSchema(graph.Config{
		Resolvers:  &Resolver{},
		Directives: CoreDirective,
	})
}

// permission for admin
var CoreDirective = graph.DirectiveRoot{
	HasRole: func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {

		//
		if !middleware.HaveToken(ctx) {
			return nil, fmt.Errorf("unauthorized")
		}

		//Get token and parse
		tokenStr := middleware.Token(ctx)
		result, err := jwt.ParseToken(tokenStr)

		if err != nil {
			return nil, fmt.Errorf("auth core: %v", err)
		}

		userRole := model.Role(result.Role) //
		roleMatch := false
		if userRole == role {
			roleMatch = true
		}

		if !roleMatch {
			// No role match found, deny access
			return nil, fmt.Errorf("access denied: you have not a permission")
		}

		ctx = context.WithValue(ctx, "username", result.Username)
		ctx = context.WithValue(ctx, "role", result.Role)
		ctx = context.WithValue(ctx, "exp", result.ExpireTime)

		fmt.Println("PASS Directive, saved user in4 to ctx!")
		return next(ctx)
	},
}
