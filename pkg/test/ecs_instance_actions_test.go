// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateEcsInstance(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryEcsInstanceFromLocal(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateEcsInstance Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No EcsInstance found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateEcsInstanceParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateEcsInstanceParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateEcsInstance(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateEcsInstance error: %v", err)
		return
	}
	golog.Infof("UpdateEcsInstance result: %s", result.UUID)
}

func TestDeleteEcsInstance(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteEcsInstance is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryEcsInstanceFromLocal(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteEcsInstance Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No EcsInstance found to test Delete")
		return
	}

	err = accountLoginCli.DeleteEcsInstance(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteEcsInstance error: %v", err)
		return
	}
	golog.Infof("DeleteEcsInstance succeeded for UUID: %s", list[0].UUID)
}

func TestStartEcsInstance(t *testing.T) {
	// Start operation - requires a stopped resource UUID
	t.Skip("TestStartEcsInstance requires a stopped resource UUID")
	// queryParam := param.NewQueryParam()
	// queryParam.AddQ("state=Stopped")
	// queryParam.Limit(1)
	// list, err := accountLoginCli.Query(context.Background(), &queryParam)
	// if err != nil || len(list) == 0 {
	// 	t.Skip("No stopped  found")
	// 	return
	// }
	// startParam := param.StartEcsInstanceParam{}
	// result, err := accountLoginCli.StartEcsInstance(context.Background(), list[0].UUID, startParam)
	// if err != nil {
	// 	t.Errorf("TestStartEcsInstance error: %v", err)
	// }
	// golog.Infof("StartEcsInstance result: %v", result)

}

func TestStopEcsInstance(t *testing.T) {
	// Stop operation - requires a running resource UUID
	t.Skip("TestStopEcsInstance requires a running resource UUID")
	// queryParam := param.NewQueryParam()
	// queryParam.AddQ("state=Running")
	// queryParam.Limit(1)
	// list, err := accountLoginCli.QueryEcsInstance(context.Background(), &queryParam)
	// if err != nil || len(list) == 0 {
	// 	t.Skip("No running EcsInstance found")
	// 	return
	// }
	// stopParam := param.StopEcsInstanceParam{}
	// result, err := accountLoginCli.StopEcsInstance(context.Background(), list[0].UUID, stopParam)
	// if err != nil {
	// 	t.Errorf("TestStopEcsInstance error: %v", err)
	// }
	// golog.Infof("StopEcsInstance result: %v", result)

}

func TestRebootEcsInstance(t *testing.T) {
	// Reboot operation - requires a running resource
	t.Skip("TestRebootEcsInstance requires a running resource UUID")

}
