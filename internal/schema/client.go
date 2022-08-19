package schema

import "github.com/graphql-go/graphql"

type ClientOptions struct {
	UserResolve func(id int) (interface{}, error)
}

type Client struct {
	userResolve func(id int) (interface{}, error)
}

func New(opts *ClientOptions) *Client {
	return &Client{
		userResolve: opts.UserResolve,
	}
}

func (c *Client) GetSchema() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: c.getQuery(),
	})
}
