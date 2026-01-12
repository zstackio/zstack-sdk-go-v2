// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVpcFirewall(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVpcFirewall(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVpcFirewall error: %v", err)
		return
	}
	golog.Infof("QueryVpcFirewall result count: %d", len(result))
}

func TestUpdateVpcFirewall(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVpcFirewall(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateVpcFirewall Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VpcFirewall found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateVpcFirewallParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateVpcFirewallParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateVpcFirewall(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateVpcFirewall error: %v", err)
		return
	}
	golog.Infof("UpdateVpcFirewall result: %s", result.Uuid)
}

func TestCreateVpcFirewall(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateVpcFirewall is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateVpcFirewallParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateVpcFirewallParamDetail{
	// 		Name: "test-vpcfirewall",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateVpcFirewall(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateVpcFirewall error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateVpcFirewall result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteVpcFirewall(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteVpcFirewall error: %v", err)
	// }
}
