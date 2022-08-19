package schema

import "github.com/graphql-go/graphql"

func (c *Client) getMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: graphql.Fields{},
	})
}
