// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVCenterPrimaryStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVCenterPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVCenterPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("QueryVCenterPrimaryStorage result count: %d", len(result))
}
func TestGetVCenterPrimaryStorage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVCenterPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestGetVCenterPrimaryStorage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VCenterPrimaryStorage found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVCenterPrimaryStorage(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVCenterPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("GetVCenterPrimaryStorage result: %s", result.UUID)
}
