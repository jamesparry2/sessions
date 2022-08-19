package datastore

type DBEngineIface interface{}

type ClientOptions struct {
	DBEngine DBEngineIface
}

type Client struct {
	dbEngine DBEngineIface
}

func New(opts *ClientOptions) *Client {
	return &Client{
		dbEngine: opts.DBEngine,
	}
}
