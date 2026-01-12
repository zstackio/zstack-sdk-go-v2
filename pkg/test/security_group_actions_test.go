// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySecurityGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySecurityGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySecurityGroup error: %v", err)
		return
	}
	golog.Infof("QuerySecurityGroup result count: %d", len(result))
}
func TestGetSecurityGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySecurityGroup(&queryParam)
	if err != nil {
		t.Errorf("TestGetSecurityGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SecurityGroup found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSecurityGroup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSecurityGroup error: %v", err)
		return
	}
	golog.Infof("GetSecurityGroup result: %s", result.UUID)
}

func TestUpdateSecurityGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySecurityGroup(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSecurityGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SecurityGroup found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSecurityGroupParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSecurityGroupParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSecurityGroup(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSecurityGroup error: %v", err)
		return
	}
	golog.Infof("UpdateSecurityGroup result: %s", result.UUID)
}

func TestDeleteSecurityGroup(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteSecurityGroup is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySecurityGroup(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteSecurityGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SecurityGroup found to test Delete")
		return
	}

	err = accountLoginCli.DeleteSecurityGroup(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteSecurityGroup error: %v", err)
		return
	}
	golog.Infof("DeleteSecurityGroup succeeded for UUID: %s", list[0].UUID)
}

func TestCreateSecurityGroup(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSecurityGroup is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSecurityGroupParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSecurityGroupParamDetail{
	// 		Name: "test-securitygroup",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSecurityGroup(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSecurityGroup error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSecurityGroup result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSecurityGroup(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSecurityGroup error: %v", err)
	// }
}
