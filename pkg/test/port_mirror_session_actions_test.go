// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPortMirrorSession(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPortMirrorSession(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPortMirrorSession error: %v", err)
		return
	}
	golog.Infof("QueryPortMirrorSession result count: %d", len(result))
}

func TestDeletePortMirrorSession(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeletePortMirrorSession is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPortMirrorSession(&queryParam)
	if err != nil {
		t.Errorf("TestDeletePortMirrorSession Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PortMirrorSession found to test Delete")
		return
	}

	err = accountLoginCli.DeletePortMirrorSession(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeletePortMirrorSession error: %v", err)
		return
	}
	golog.Infof("DeletePortMirrorSession succeeded for UUID: %s", list[0].UUID)
}

func TestCreatePortMirrorSession(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreatePortMirrorSession is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreatePortMirrorSessionParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreatePortMirrorSessionParamDetail{
	// 		Name: "test-portmirrorsession",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreatePortMirrorSession(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreatePortMirrorSession error: %v", err)
	// 	return
	// }
	// golog.Infof("CreatePortMirrorSession result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeletePortMirrorSession(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeletePortMirrorSession error: %v", err)
	// }
}
