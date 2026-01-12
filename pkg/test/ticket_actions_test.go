// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryTicket(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryTicket(&queryParam)
	if err != nil {
		t.Errorf("TestQueryTicket error: %v", err)
		return
	}
	golog.Infof("QueryTicket result count: %d", len(result))
}

func TestDeleteTicket(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteTicket is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryTicket(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteTicket Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Ticket found to test Delete")
		return
	}

	err = accountLoginCli.DeleteTicket(list[0].Uuid, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteTicket error: %v", err)
		return
	}
	golog.Infof("DeleteTicket succeeded for UUID: %s", list[0].Uuid)
}

func TestCreateTicket(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateTicket is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateTicketParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateTicketParamDetail{
	// 		Name: "test-ticket",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateTicket(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateTicket error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateTicket result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteTicket(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteTicket error: %v", err)
	// }
}
