// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryImageReplicationGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryImageReplicationGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryImageReplicationGroup error: %v", err)
		return
	}
	golog.Infof("QueryImageReplicationGroup result count: %d", len(result))
}

func TestDeleteImageReplicationGroup(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteImageReplicationGroup is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryImageReplicationGroup(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteImageReplicationGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ImageReplicationGroup found to test Delete")
		return
	}

	err = accountLoginCli.DeleteImageReplicationGroup(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteImageReplicationGroup error: %v", err)
		return
	}
	golog.Infof("DeleteImageReplicationGroup succeeded for UUID: %s", list[0].UUID)
}

func TestCreateImageReplicationGroup(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateImageReplicationGroup is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateImageReplicationGroupParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateImageReplicationGroupParamDetail{
	// 		Name: "test-imagereplicationgroup",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateImageReplicationGroup(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateImageReplicationGroup error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateImageReplicationGroup result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteImageReplicationGroup(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteImageReplicationGroup error: %v", err)
	// }
}
