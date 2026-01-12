// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPolicyRouteRule(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPolicyRouteRule(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPolicyRouteRule error: %v", err)
		return
	}
	golog.Infof("QueryPolicyRouteRule result count: %d", len(result))
}

func TestDeletePolicyRouteRule(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeletePolicyRouteRule is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPolicyRouteRule(&queryParam)
	if err != nil {
		t.Errorf("TestDeletePolicyRouteRule Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PolicyRouteRule found to test Delete")
		return
	}

	err = accountLoginCli.DeletePolicyRouteRule(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeletePolicyRouteRule error: %v", err)
		return
	}
	golog.Infof("DeletePolicyRouteRule succeeded for UUID: %s", list[0].UUID)
}

func TestCreatePolicyRouteRule(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreatePolicyRouteRule is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreatePolicyRouteRuleParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreatePolicyRouteRuleParamDetail{
	// 		Name: "test-policyrouterule",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreatePolicyRouteRule(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreatePolicyRouteRule error: %v", err)
	// 	return
	// }
	// golog.Infof("CreatePolicyRouteRule result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeletePolicyRouteRule(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeletePolicyRouteRule error: %v", err)
	// }
}
