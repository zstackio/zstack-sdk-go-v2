// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVCenterBackupStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVCenterBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVCenterBackupStorage error: %v", err)
		return
	}
	golog.Infof("QueryVCenterBackupStorage result count: %d", len(result))
}
func TestGetVCenterBackupStorage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVCenterBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestGetVCenterBackupStorage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VCenterBackupStorage found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVCenterBackupStorage(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVCenterBackupStorage error: %v", err)
		return
	}
	golog.Infof("GetVCenterBackupStorage result: %s", result.UUID)
}
