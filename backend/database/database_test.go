package database_test

import (
	"fmt"
	"testing"

	"github.com/supabase-community/supabase-go"
)

const (
	API_URL = ""
	API_KEY = ""
)

func TestFrom(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	data, count, err := client.From("account").Select("*", "exact", false).Execute()
	fmt.Println(string(data), err, count)
}

func TestRpc(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	result := client.Rpc("hello_world", "", nil)
	fmt.Println(result)
}
