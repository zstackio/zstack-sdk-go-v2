// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNvmeLun(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryNvmeLun(&queryParam)
	if err != nil {
		t.Errorf("TestQueryNvmeLun error: %v", err)
		return
	}
	golog.Infof("QueryNvmeLun result count: %d", len(result))
}
func TestGetNvmeLun(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryNvmeLun(&queryParam)
	if err != nil {
		t.Errorf("TestGetNvmeLun Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No NvmeLun found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetNvmeLun(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetNvmeLun error: %v", err)
		return
	}
	golog.Infof("GetNvmeLun result: %s", result.UUID)
}
