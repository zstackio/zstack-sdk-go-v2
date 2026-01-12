// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHuaweiIMasterVRouter(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryHuaweiIMasterVRouter(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHuaweiIMasterVRouter error: %v", err)
		return
	}
	golog.Infof("QueryHuaweiIMasterVRouter result count: %d", len(result))
}
func TestGetHuaweiIMasterVRouter(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryHuaweiIMasterVRouter(&queryParam)
	if err != nil {
		t.Errorf("TestGetHuaweiIMasterVRouter Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No HuaweiIMasterVRouter found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetHuaweiIMasterVRouter(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetHuaweiIMasterVRouter error: %v", err)
		return
	}
	golog.Infof("GetHuaweiIMasterVRouter result: %s", result.UUID)
}

func TestDeleteHuaweiIMasterVRouter(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteHuaweiIMasterVRouter is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryHuaweiIMasterVRouter(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteHuaweiIMasterVRouter Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No HuaweiIMasterVRouter found to test Delete")
		return
	}

	err = accountLoginCli.DeleteHuaweiIMasterVRouter(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteHuaweiIMasterVRouter error: %v", err)
		return
	}
	golog.Infof("DeleteHuaweiIMasterVRouter succeeded for UUID: %s", list[0].UUID)
}

func TestCreateHuaweiIMasterVRouter(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateHuaweiIMasterVRouter is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateHuaweiIMasterVRouterParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateHuaweiIMasterVRouterParamDetail{
	// 		Name: "test-huaweiimastervrouter",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateHuaweiIMasterVRouter(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateHuaweiIMasterVRouter error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateHuaweiIMasterVRouter result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteHuaweiIMasterVRouter(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteHuaweiIMasterVRouter error: %v", err)
	// }
}
