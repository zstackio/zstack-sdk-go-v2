// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryOvnController(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryOvnController(&queryParam)
	if err != nil {
		t.Errorf("TestQueryOvnController error: %v", err)
		return
	}
	golog.Infof("QueryOvnController result count: %d", len(result))
}
func TestGetOvnController(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryOvnController(&queryParam)
	if err != nil {
		t.Errorf("TestGetOvnController Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No OvnController found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetOvnController(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetOvnController error: %v", err)
		return
	}
	golog.Infof("GetOvnController result: %s", result.UUID)
}
