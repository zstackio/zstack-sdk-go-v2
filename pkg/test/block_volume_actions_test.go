// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBlockVolume(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBlockVolume(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBlockVolume error: %v", err)
		return
	}
	golog.Infof("QueryBlockVolume result count: %d", len(result))
}

func TestUpdateBlockVolume(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBlockVolume(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateBlockVolume Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BlockVolume found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateBlockVolumeParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateBlockVolumeParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateBlockVolume(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateBlockVolume error: %v", err)
		return
	}
	golog.Infof("UpdateBlockVolume result: %s", result.UUID)
}

func TestCreateBlockVolume(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateBlockVolume is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateBlockVolumeParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateBlockVolumeParamDetail{
	// 		Name: "test-blockvolume",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateBlockVolume(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateBlockVolume error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateBlockVolume result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteBlockVolume(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteBlockVolume error: %v", err)
	// }
}
