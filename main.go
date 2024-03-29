package main

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"latestPost": &graphql.Field{
			Type: graphql.String,
			// Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 	return "Hello World", nil
			// },
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	},
})

// Schema ...d
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

func main() {

	// create a graphl-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := handler.New(&handler.Config{
		Schema:   &Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)

	// and serve!
	http.ListenAndServe(":8080", nil)
}
