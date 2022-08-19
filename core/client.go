package core

import "github.com/jamesparry2/sessions/internal/datastore"

type ClientOptions struct {
	DBClient datastore.DBEngineIface
}

type Client struct {
	dbClient datastore.DBEngineIface
}

func New(opts *ClientOptions) *Client {
	return &Client{
		dbClient: opts.DBClient,
	}
}
