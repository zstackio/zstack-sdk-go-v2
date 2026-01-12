// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIscsiLun(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryIscsiLun(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIscsiLun error: %v", err)
		return
	}
	golog.Infof("QueryIscsiLun result count: %d", len(result))
}
func TestGetIscsiLun(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIscsiLun(&queryParam)
	if err != nil {
		t.Errorf("TestGetIscsiLun Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IscsiLun found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetIscsiLun(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetIscsiLun error: %v", err)
		return
	}
	golog.Infof("GetIscsiLun result: %s", result.UUID)
}
