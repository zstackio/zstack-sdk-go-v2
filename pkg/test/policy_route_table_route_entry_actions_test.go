// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPolicyRouteTableRouteEntry(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPolicyRouteTableRouteEntry(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPolicyRouteTableRouteEntry error: %v", err)
		return
	}
	golog.Infof("QueryPolicyRouteTableRouteEntry result count: %d", len(result))
}
func TestGetPolicyRouteTableRouteEntry(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPolicyRouteTableRouteEntry(&queryParam)
	if err != nil {
		t.Errorf("TestGetPolicyRouteTableRouteEntry Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PolicyRouteTableRouteEntry found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetPolicyRouteTableRouteEntry(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetPolicyRouteTableRouteEntry error: %v", err)
		return
	}
	golog.Infof("GetPolicyRouteTableRouteEntry result: %s", result.UUID)
}

func TestDeletePolicyRouteTableRouteEntry(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeletePolicyRouteTableRouteEntry is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPolicyRouteTableRouteEntry(&queryParam)
	if err != nil {
		t.Errorf("TestDeletePolicyRouteTableRouteEntry Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PolicyRouteTableRouteEntry found to test Delete")
		return
	}

	err = accountLoginCli.DeletePolicyRouteTableRouteEntry(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeletePolicyRouteTableRouteEntry error: %v", err)
		return
	}
	golog.Infof("DeletePolicyRouteTableRouteEntry succeeded for UUID: %s", list[0].UUID)
}

func TestCreatePolicyRouteTableRouteEntry(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreatePolicyRouteTableRouteEntry is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreatePolicyRouteTableRouteEntryParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreatePolicyRouteTableRouteEntryParamDetail{
	// 		Name: "test-policyroutetablerouteentry",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreatePolicyRouteTableRouteEntry(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreatePolicyRouteTableRouteEntry error: %v", err)
	// 	return
	// }
	// golog.Infof("CreatePolicyRouteTableRouteEntry result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeletePolicyRouteTableRouteEntry(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeletePolicyRouteTableRouteEntry error: %v", err)
	// }
}
