// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryTicketType(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryTicketType(&queryParam)
	if err != nil {
		t.Errorf("TestQueryTicketType error: %v", err)
		return
	}
	golog.Infof("QueryTicketType result count: %d", len(result))
}
func TestGetTicketType(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryTicketType(&queryParam)
	if err != nil {
		t.Errorf("TestGetTicketType Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No TicketType found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetTicketType(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetTicketType error: %v", err)
		return
	}
	golog.Infof("GetTicketType result: %s", result.UUID)
}
