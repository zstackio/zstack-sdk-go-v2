// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPortForwardingRule(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPortForwardingRule(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPortForwardingRule error: %v", err)
		return
	}
	golog.Infof("QueryPortForwardingRule result count: %d", len(result))
}

func TestUpdatePortForwardingRule(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPortForwardingRule(&queryParam)
	if err != nil {
		t.Errorf("TestUpdatePortForwardingRule Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PortForwardingRule found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdatePortForwardingRuleParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdatePortForwardingRuleParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdatePortForwardingRule(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdatePortForwardingRule error: %v", err)
		return
	}
	golog.Infof("UpdatePortForwardingRule result: %s", result.UUID)
}

func TestDeletePortForwardingRule(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeletePortForwardingRule is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPortForwardingRule(&queryParam)
	if err != nil {
		t.Errorf("TestDeletePortForwardingRule Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PortForwardingRule found to test Delete")
		return
	}

	err = accountLoginCli.DeletePortForwardingRule(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeletePortForwardingRule error: %v", err)
		return
	}
	golog.Infof("DeletePortForwardingRule succeeded for UUID: %s", list[0].UUID)
}

func TestCreatePortForwardingRule(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreatePortForwardingRule is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreatePortForwardingRuleParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreatePortForwardingRuleParamDetail{
	// 		Name: "test-portforwardingrule",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreatePortForwardingRule(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreatePortForwardingRule error: %v", err)
	// 	return
	// }
	// golog.Infof("CreatePortForwardingRule result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeletePortForwardingRule(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeletePortForwardingRule error: %v", err)
	// }
}

func TestDetachPortForwardingRule(t *testing.T) {
	// Detach operation
	t.Skip("TestDetachPortForwardingRule requires an attached resource")

}

func TestAttachPortForwardingRule(t *testing.T) {
	// Attach operation
	t.Skip("TestAttachPortForwardingRule requires valid resource UUIDs to attach")

}
