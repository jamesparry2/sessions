package schema

import "github.com/graphql-go/graphql"

func (c *Client) getQuery() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"user": &graphql.Field{
					Type:        c.getUserObject(),
					Description: "Get user by ID",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						s, _ := p.Args["id"].(int)
						return c.userResolve(s)
					},
				},
			},
		},
	)
}
