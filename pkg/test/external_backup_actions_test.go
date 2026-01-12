// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryExternalBackup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryExternalBackup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryExternalBackup error: %v", err)
		return
	}
	golog.Infof("QueryExternalBackup result count: %d", len(result))
}
func TestGetExternalBackup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryExternalBackup(&queryParam)
	if err != nil {
		t.Errorf("TestGetExternalBackup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ExternalBackup found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetExternalBackup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetExternalBackup error: %v", err)
		return
	}
	golog.Infof("GetExternalBackup result: %s", result.UUID)
}

func TestDeleteExternalBackup(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteExternalBackup is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryExternalBackup(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteExternalBackup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ExternalBackup found to test Delete")
		return
	}

	err = accountLoginCli.DeleteExternalBackup(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteExternalBackup error: %v", err)
		return
	}
	golog.Infof("DeleteExternalBackup succeeded for UUID: %s", list[0].UUID)
}
