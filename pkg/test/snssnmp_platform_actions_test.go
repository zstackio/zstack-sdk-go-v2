// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSSnmpPlatform(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSSnmpPlatform(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSSnmpPlatform error: %v", err)
		return
	}
	golog.Infof("QuerySNSSnmpPlatform result count: %d", len(result))
}

func TestUpdateSNSSnmpPlatform(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSSnmpPlatform(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSNSSnmpPlatform Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSSnmpPlatform found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSNSSnmpPlatformParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSNSSnmpPlatformParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSNSSnmpPlatform(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSNSSnmpPlatform error: %v", err)
		return
	}
	golog.Infof("UpdateSNSSnmpPlatform result: %s", result.Uuid)
}

func TestCreateSNSSnmpPlatform(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSNSSnmpPlatform is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSNSSnmpPlatformParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSNSSnmpPlatformParamDetail{
	// 		Name: "test-snssnmpplatform",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSNSSnmpPlatform(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSNSSnmpPlatform error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSNSSnmpPlatform result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSNSSnmpPlatform(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSNSSnmpPlatform error: %v", err)
	// }
}
