package rpc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/infraboard/eventbox/apps/book"
	client "github.com/infraboard/eventbox/client/rpc"
	"github.com/infraboard/mcenter/client/rpc"
)

func TestBookQuery(t *testing.T) {
	should := assert.New(t)

	conf := rpc.NewDefaultConfig()
	c, err := client.NewClient(conf)
	if should.NoError(err) {
		resp, err := c.Book().QueryBook(
			context.Background(),
			book.NewQueryBookRequest(),
		)
		should.NoError(err)
		fmt.Println(resp.Items)
	}
}
