package handler

import (
	"github.com/graphql-go/graphql"
	h "github.com/graphql-go/handler"
)

type ClientOptions struct {
	GraphExecute func(graphql.Params) *graphql.Result
}

type Client struct {
	graphExecute func(graphql.Params) *graphql.Result
}

func New(opts *ClientOptions) *Client {
	return &Client{
		graphExecute: opts.GraphExecute,
	}
}

func (c *Client) Handler(schema graphql.Schema) *h.Handler {
	return h.New(&h.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: false,
	})

}
