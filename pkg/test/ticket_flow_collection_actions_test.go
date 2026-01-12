// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryTicketFlowCollection(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryTicketFlowCollection(&queryParam)
	if err != nil {
		t.Errorf("TestQueryTicketFlowCollection error: %v", err)
		return
	}
	golog.Infof("QueryTicketFlowCollection result count: %d", len(result))
}

func TestDeleteTicketFlowCollection(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteTicketFlowCollection is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryTicketFlowCollection(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteTicketFlowCollection Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No TicketFlowCollection found to test Delete")
		return
	}

	err = accountLoginCli.DeleteTicketFlowCollection(list[0].Uuid, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteTicketFlowCollection error: %v", err)
		return
	}
	golog.Infof("DeleteTicketFlowCollection succeeded for UUID: %s", list[0].Uuid)
}
