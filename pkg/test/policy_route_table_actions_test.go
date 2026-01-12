// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPolicyRouteTable(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPolicyRouteTable(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPolicyRouteTable error: %v", err)
		return
	}
	golog.Infof("QueryPolicyRouteTable result count: %d", len(result))
}

func TestDeletePolicyRouteTable(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeletePolicyRouteTable is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPolicyRouteTable(&queryParam)
	if err != nil {
		t.Errorf("TestDeletePolicyRouteTable Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PolicyRouteTable found to test Delete")
		return
	}

	err = accountLoginCli.DeletePolicyRouteTable(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeletePolicyRouteTable error: %v", err)
		return
	}
	golog.Infof("DeletePolicyRouteTable succeeded for UUID: %s", list[0].UUID)
}

func TestCreatePolicyRouteTable(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreatePolicyRouteTable is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreatePolicyRouteTableParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreatePolicyRouteTableParamDetail{
	// 		Name: "test-policyroutetable",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreatePolicyRouteTable(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreatePolicyRouteTable error: %v", err)
	// 	return
	// }
	// golog.Infof("CreatePolicyRouteTable result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeletePolicyRouteTable(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeletePolicyRouteTable error: %v", err)
	// }
}
