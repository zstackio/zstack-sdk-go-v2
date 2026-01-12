// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryLoadBalancerServerGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryLoadBalancerServerGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryLoadBalancerServerGroup error: %v", err)
		return
	}
	golog.Infof("QueryLoadBalancerServerGroup result count: %d", len(result))
}
func TestGetLoadBalancerServerGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryLoadBalancerServerGroup(&queryParam)
	if err != nil {
		t.Errorf("TestGetLoadBalancerServerGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LoadBalancerServerGroup found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetLoadBalancerServerGroup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetLoadBalancerServerGroup error: %v", err)
		return
	}
	golog.Infof("GetLoadBalancerServerGroup result: %s", result.UUID)
}

func TestUpdateLoadBalancerServerGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryLoadBalancerServerGroup(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateLoadBalancerServerGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LoadBalancerServerGroup found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateLoadBalancerServerGroupParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateLoadBalancerServerGroupParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateLoadBalancerServerGroup(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateLoadBalancerServerGroup error: %v", err)
		return
	}
	golog.Infof("UpdateLoadBalancerServerGroup result: %s", result.UUID)
}

func TestDeleteLoadBalancerServerGroup(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteLoadBalancerServerGroup is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryLoadBalancerServerGroup(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteLoadBalancerServerGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LoadBalancerServerGroup found to test Delete")
		return
	}

	err = accountLoginCli.DeleteLoadBalancerServerGroup(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteLoadBalancerServerGroup error: %v", err)
		return
	}
	golog.Infof("DeleteLoadBalancerServerGroup succeeded for UUID: %s", list[0].UUID)
}

func TestCreateLoadBalancerServerGroup(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateLoadBalancerServerGroup is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateLoadBalancerServerGroupParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateLoadBalancerServerGroupParamDetail{
	// 		Name: "test-loadbalancerservergroup",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateLoadBalancerServerGroup(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateLoadBalancerServerGroup error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateLoadBalancerServerGroup result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteLoadBalancerServerGroup(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteLoadBalancerServerGroup error: %v", err)
	// }
}
