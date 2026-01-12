// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySharedBlockGroupPrimaryStorageHostRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySharedBlockGroupPrimaryStorageHostRef(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySharedBlockGroupPrimaryStorageHostRef error: %v", err)
		return
	}
	golog.Infof("QuerySharedBlockGroupPrimaryStorageHostRef result count: %d", len(result))
}
func TestGetSharedBlockGroupPrimaryStorageHostRef(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySharedBlockGroupPrimaryStorageHostRef(&queryParam)
	if err != nil {
		t.Errorf("TestGetSharedBlockGroupPrimaryStorageHostRef Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SharedBlockGroupPrimaryStorageHostRef found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSharedBlockGroupPrimaryStorageHostRef(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSharedBlockGroupPrimaryStorageHostRef error: %v", err)
		return
	}
	golog.Infof("GetSharedBlockGroupPrimaryStorageHostRef result: %s", result.UUID)
}
