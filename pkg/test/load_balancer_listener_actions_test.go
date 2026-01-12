// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryLoadBalancerListener(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryLoadBalancerListener(&queryParam)
	if err != nil {
		t.Errorf("TestQueryLoadBalancerListener error: %v", err)
		return
	}
	golog.Infof("QueryLoadBalancerListener result count: %d", len(result))
}
func TestGetLoadBalancerListener(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryLoadBalancerListener(&queryParam)
	if err != nil {
		t.Errorf("TestGetLoadBalancerListener Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LoadBalancerListener found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetLoadBalancerListener(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetLoadBalancerListener error: %v", err)
		return
	}
	golog.Infof("GetLoadBalancerListener result: %s", result.UUID)
}

func TestUpdateLoadBalancerListener(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryLoadBalancerListener(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateLoadBalancerListener Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LoadBalancerListener found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateLoadBalancerListenerParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateLoadBalancerListenerParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateLoadBalancerListener(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateLoadBalancerListener error: %v", err)
		return
	}
	golog.Infof("UpdateLoadBalancerListener result: %s", result.UUID)
}

func TestDeleteLoadBalancerListener(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteLoadBalancerListener is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryLoadBalancerListener(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteLoadBalancerListener Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LoadBalancerListener found to test Delete")
		return
	}

	err = accountLoginCli.DeleteLoadBalancerListener(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteLoadBalancerListener error: %v", err)
		return
	}
	golog.Infof("DeleteLoadBalancerListener succeeded for UUID: %s", list[0].UUID)
}

func TestCreateLoadBalancerListener(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateLoadBalancerListener is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateLoadBalancerListenerParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateLoadBalancerListenerParamDetail{
	// 		Name: "test-loadbalancerlistener",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateLoadBalancerListener(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateLoadBalancerListener error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateLoadBalancerListener result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteLoadBalancerListener(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteLoadBalancerListener error: %v", err)
	// }
}

func TestChangeLoadBalancerListener(t *testing.T) {
	// Change operation
	t.Skip("TestChangeLoadBalancerListener requires specific parameters")

}
