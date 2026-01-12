// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVolume(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVolume(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVolume error: %v", err)
		return
	}
	golog.Infof("QueryVolume result count: %d", len(result))
}
func TestGetVolume(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVolume(&queryParam)
	if err != nil {
		t.Errorf("TestGetVolume Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Volume found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVolume(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVolume error: %v", err)
		return
	}
	golog.Infof("GetVolume result: %s", result.UUID)
}

func TestUpdateVolume(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVolume(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateVolume Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Volume found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateVolumeParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateVolumeParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateVolume(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateVolume error: %v", err)
		return
	}
	golog.Infof("UpdateVolume result: %s", result.UUID)
}
