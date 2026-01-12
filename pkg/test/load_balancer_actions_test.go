// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryLoadBalancer(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryLoadBalancer(&queryParam)
	if err != nil {
		t.Errorf("TestQueryLoadBalancer error: %v", err)
		return
	}
	golog.Infof("QueryLoadBalancer result count: %d", len(result))
}

func TestUpdateLoadBalancer(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryLoadBalancer(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateLoadBalancer Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LoadBalancer found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateLoadBalancerParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateLoadBalancerParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateLoadBalancer(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateLoadBalancer error: %v", err)
		return
	}
	golog.Infof("UpdateLoadBalancer result: %s", result.UUID)
}

func TestDeleteLoadBalancer(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteLoadBalancer is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryLoadBalancer(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteLoadBalancer Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LoadBalancer found to test Delete")
		return
	}

	err = accountLoginCli.DeleteLoadBalancer(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteLoadBalancer error: %v", err)
		return
	}
	golog.Infof("DeleteLoadBalancer succeeded for UUID: %s", list[0].UUID)
}

func TestCreateLoadBalancer(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateLoadBalancer is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateLoadBalancerParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateLoadBalancerParamDetail{
	// 		Name: "test-loadbalancer",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateLoadBalancer(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateLoadBalancer error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateLoadBalancer result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteLoadBalancer(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteLoadBalancer error: %v", err)
	// }
}

func TestRefreshLoadBalancer(t *testing.T) {
	// RefreshLoadBalancer operation
	t.Skip("TestRefreshLoadBalancer requires manual implementation")

}
