package database

import (
	"errors"
	"os"

	"github.com/supabase/postgrest-go"
)

const REST_URL = "/rest/v1"

type Client struct {
	rest *postgrest.Client
}

type RestOptions struct {
	Schema string
}

type ClientOptions struct {
	Headers map[string]string
	Db      *RestOptions
}

// NewClient creates a new Postgrest client.
func NewClient(bearerKey string, options *ClientOptions, schema string) (*Client, error) {
	url := os.Getenv("API_URL")
	key := os.Getenv("API_KEY")

	// Remove "Bearer "
	bearerKey = bearerKey[7:]

	if url == "" || key == "" {
		return nil, errors.New("url and key are required")
	}

	defaultHeaders := map[string]string{
		"Authorization": "Bearer " + bearerKey,
		"apikey":        key,
	}

	client := &Client{}

	if options != nil && options.Db != nil {
		if len(options.Headers) > 0 {
			for k, v := range options.Headers {
				defaultHeaders[k] = v
			}
		}
	}
	client.rest = postgrest.NewClient(url+REST_URL, schema, defaultHeaders)

	return client, nil
}

// From returns a QueryBuilder for the specified table.
func (c *Client) From(table string) *postgrest.QueryBuilder {
	return c.rest.From(table)
}

// Rpc returns a string for the specified function.
func (c *Client) Rpc(name, count string, rpcBody interface{}) string {
	return c.rest.Rpc(name, count, rpcBody)
}
