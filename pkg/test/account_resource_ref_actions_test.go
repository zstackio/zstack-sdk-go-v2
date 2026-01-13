// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAccountResourceRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAccountResourceRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAccountResourceRef error: %v", err)
		return
	}
	golog.Infof("QueryAccountResourceRef result count: %d", len(result))
}

func TestGetAccountResourceRef(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAccountResourceRef(&queryParam)
	if err != nil {
		t.Errorf("TestGetAccountResourceRef Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AccountResourceRef found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetAccountResourceRef(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetAccountResourceRef error: %v", err)
		return
	}
	golog.Infof("GetAccountResourceRef result: %s", result.UUID)
}
