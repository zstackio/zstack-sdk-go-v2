// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryTicketFlow(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryTicketFlow(&queryParam)
	if err != nil {
		t.Errorf("TestQueryTicketFlow error: %v", err)
		return
	}
	golog.Infof("QueryTicketFlow result count: %d", len(result))
}
func TestGetTicketFlow(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryTicketFlow(&queryParam)
	if err != nil {
		t.Errorf("TestGetTicketFlow Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No TicketFlow found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetTicketFlow(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetTicketFlow error: %v", err)
		return
	}
	golog.Infof("GetTicketFlow result: %s", result.UUID)
}
