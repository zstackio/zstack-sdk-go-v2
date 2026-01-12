// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySlbVmInstance(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySlbVmInstance(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySlbVmInstance error: %v", err)
		return
	}
	golog.Infof("QuerySlbVmInstance result count: %d", len(result))
}
func TestGetSlbVmInstance(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySlbVmInstance(&queryParam)
	if err != nil {
		t.Errorf("TestGetSlbVmInstance Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SlbVmInstance found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSlbVmInstance(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSlbVmInstance error: %v", err)
		return
	}
	golog.Infof("GetSlbVmInstance result: %s", result.UUID)
}
