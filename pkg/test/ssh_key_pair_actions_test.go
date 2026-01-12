// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySshKeyPair(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySshKeyPair(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySshKeyPair error: %v", err)
		return
	}
	golog.Infof("QuerySshKeyPair result count: %d", len(result))
}
func TestGetSshKeyPair(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySshKeyPair(&queryParam)
	if err != nil {
		t.Errorf("TestGetSshKeyPair Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SshKeyPair found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSshKeyPair(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSshKeyPair error: %v", err)
		return
	}
	golog.Infof("GetSshKeyPair result: %s", result.UUID)
}

func TestUpdateSshKeyPair(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySshKeyPair(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSshKeyPair Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SshKeyPair found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSshKeyPairParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSshKeyPairParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSshKeyPair(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSshKeyPair error: %v", err)
		return
	}
	golog.Infof("UpdateSshKeyPair result: %s", result.UUID)
}

func TestDeleteSshKeyPair(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteSshKeyPair is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySshKeyPair(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteSshKeyPair Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SshKeyPair found to test Delete")
		return
	}

	err = accountLoginCli.DeleteSshKeyPair(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteSshKeyPair error: %v", err)
		return
	}
	golog.Infof("DeleteSshKeyPair succeeded for UUID: %s", list[0].UUID)
}

func TestCreateSshKeyPair(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSshKeyPair is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSshKeyPairParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSshKeyPairParamDetail{
	// 		Name: "test-sshkeypair",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSshKeyPair(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSshKeyPair error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSshKeyPair result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSshKeyPair(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSshKeyPair error: %v", err)
	// }
}

func TestGenerateSshKeyPair(t *testing.T) {
	// GenerateSshKeyPair operation
	t.Skip("TestGenerateSshKeyPair requires manual implementation")

}
