package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Define a Book type
var bookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// Define a Query type
var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"book": &graphql.Field{
			Type: bookType,
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// Your data fetching logic here
				// For simplicity, return a hardcoded book
				return map[string]interface{}{
					"title":  "Sample Book",
					"author": "Sample Author",
				}, nil
			},
		},
	},
})

func main() {
	// Create a schema, passing in the root query type
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	if err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	// Create a GraphQL handler
	handler := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", handler)

	fmt.Println("GraphQL server is running on :8080")
	err = http.ListenAndServe(":8081", nil)

	if err != nil {
		log.Fatal(err)
	}
}
