// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySecretResourcePool(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySecretResourcePool(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySecretResourcePool error: %v", err)
		return
	}
	golog.Infof("QuerySecretResourcePool result count: %d", len(result))
}
func TestGetSecretResourcePool(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySecretResourcePool(&queryParam)
	if err != nil {
		t.Errorf("TestGetSecretResourcePool Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SecretResourcePool found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSecretResourcePool(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSecretResourcePool error: %v", err)
		return
	}
	golog.Infof("GetSecretResourcePool result: %s", result.UUID)
}

func TestUpdateSecretResourcePool(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySecretResourcePool(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSecretResourcePool Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SecretResourcePool found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSecretResourcePoolParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSecretResourcePoolParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSecretResourcePool(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSecretResourcePool error: %v", err)
		return
	}
	golog.Infof("UpdateSecretResourcePool result: %s", result.UUID)
}

func TestDeleteSecretResourcePool(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteSecretResourcePool is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySecretResourcePool(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteSecretResourcePool Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SecretResourcePool found to test Delete")
		return
	}

	err = accountLoginCli.DeleteSecretResourcePool(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteSecretResourcePool error: %v", err)
		return
	}
	golog.Infof("DeleteSecretResourcePool succeeded for UUID: %s", list[0].UUID)
}
