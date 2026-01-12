// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryArchiveTicket(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryArchiveTicket(&queryParam)
	if err != nil {
		t.Errorf("TestQueryArchiveTicket error: %v", err)
		return
	}
	golog.Infof("QueryArchiveTicket result count: %d", len(result))
}
func TestGetArchiveTicket(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryArchiveTicket(&queryParam)
	if err != nil {
		t.Errorf("TestGetArchiveTicket Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ArchiveTicket found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetArchiveTicket(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetArchiveTicket error: %v", err)
		return
	}
	golog.Infof("GetArchiveTicket result: %s", result.UUID)
}
