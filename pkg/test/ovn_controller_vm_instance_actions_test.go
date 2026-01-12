// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryOvnControllerVmInstance(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryOvnControllerVmInstance(&queryParam)
	if err != nil {
		t.Errorf("TestQueryOvnControllerVmInstance error: %v", err)
		return
	}
	golog.Infof("QueryOvnControllerVmInstance result count: %d", len(result))
}
func TestGetOvnControllerVmInstance(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryOvnControllerVmInstance(&queryParam)
	if err != nil {
		t.Errorf("TestGetOvnControllerVmInstance Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No OvnControllerVmInstance found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetOvnControllerVmInstance(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetOvnControllerVmInstance error: %v", err)
		return
	}
	golog.Infof("GetOvnControllerVmInstance result: %s", result.UUID)
}
