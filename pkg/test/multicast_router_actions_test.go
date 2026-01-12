// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMulticastRouter(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMulticastRouter(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMulticastRouter error: %v", err)
		return
	}
	golog.Infof("QueryMulticastRouter result count: %d", len(result))
}

func TestDeleteMulticastRouter(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteMulticastRouter is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMulticastRouter(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteMulticastRouter Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MulticastRouter found to test Delete")
		return
	}

	err = accountLoginCli.DeleteMulticastRouter(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteMulticastRouter error: %v", err)
		return
	}
	golog.Infof("DeleteMulticastRouter succeeded for UUID: %s", list[0].UUID)
}

func TestCreateMulticastRouter(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateMulticastRouter is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateMulticastRouterParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateMulticastRouterParamDetail{
	// 		Name: "test-multicastrouter",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateMulticastRouter(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateMulticastRouter error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateMulticastRouter result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteMulticastRouter(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteMulticastRouter error: %v", err)
	// }
}
