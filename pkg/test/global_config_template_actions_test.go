// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryGlobalConfigTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryGlobalConfigTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryGlobalConfigTemplate error: %v", err)
		return
	}
	golog.Infof("QueryGlobalConfigTemplate result count: %d", len(result))
}
func TestGetGlobalConfigTemplate(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryGlobalConfigTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestGetGlobalConfigTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No GlobalConfigTemplate found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetGlobalConfigTemplate(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetGlobalConfigTemplate error: %v", err)
		return
	}
	golog.Infof("GetGlobalConfigTemplate result: %s", result.UUID)
}
