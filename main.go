package main

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/jamesparry2/sessions/handler"
	"github.com/jamesparry2/sessions/internal/schema"
)

func main() {
	schemaClient := schema.New(&schema.ClientOptions{
		UserResolve: func(id int) (interface{}, error) {
			return nil, nil
		},
	})

	handlerClient := handler.New(&handler.ClientOptions{
		GraphExecute: graphql.Do,
	})

	if schema, err := schemaClient.GetSchema(); err != nil {
		fmt.Printf("failed to create a valid schema '%v'", err.Error())
	} else {
		http.Handle("/graphql", handlerClient.Handler(schema))

		fmt.Println("Server is running on port 8080")
		http.ListenAndServe(":8080", nil)
	}
}
