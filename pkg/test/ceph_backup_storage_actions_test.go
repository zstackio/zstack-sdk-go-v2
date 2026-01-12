// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCephBackupStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryCephBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCephBackupStorage error: %v", err)
		return
	}
	golog.Infof("QueryCephBackupStorage result count: %d", len(result))
}
func TestGetCephBackupStorage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCephBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestGetCephBackupStorage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No CephBackupStorage found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetCephBackupStorage(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetCephBackupStorage error: %v", err)
		return
	}
	golog.Infof("GetCephBackupStorage result: %s", result.UUID)
}

func TestAddCephBackupStorage(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddCephBackupStorage requires valid creation parameters")

}
