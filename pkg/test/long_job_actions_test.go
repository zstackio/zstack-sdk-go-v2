// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryLongJob(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryLongJob(&queryParam)
	if err != nil {
		t.Errorf("TestQueryLongJob error: %v", err)
		return
	}
	golog.Infof("QueryLongJob result count: %d", len(result))
}

func TestUpdateLongJob(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryLongJob(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateLongJob Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LongJob found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateLongJobParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateLongJobParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateLongJob(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateLongJob error: %v", err)
		return
	}
	golog.Infof("UpdateLongJob result: %s", result.UUID)
}

func TestDeleteLongJob(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteLongJob is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryLongJob(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteLongJob Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LongJob found to test Delete")
		return
	}

	err = accountLoginCli.DeleteLongJob(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteLongJob error: %v", err)
		return
	}
	golog.Infof("DeleteLongJob succeeded for UUID: %s", list[0].UUID)
}

func TestCleanLongJob(t *testing.T) {
	// CleanLongJob operation
	t.Skip("TestCleanLongJob requires manual implementation")

}

func TestResumeLongJob(t *testing.T) {
	// ResumeLongJob operation
	t.Skip("TestResumeLongJob requires manual implementation")

}
