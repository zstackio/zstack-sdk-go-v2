// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryModelEvaluationTask(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryModelEvaluationTask(&queryParam)
	if err != nil {
		t.Errorf("TestQueryModelEvaluationTask error: %v", err)
		return
	}
	golog.Infof("QueryModelEvaluationTask result count: %d", len(result))
}
func TestGetModelEvaluationTask(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryModelEvaluationTask(&queryParam)
	if err != nil {
		t.Errorf("TestGetModelEvaluationTask Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ModelEvaluationTask found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetModelEvaluationTask(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetModelEvaluationTask error: %v", err)
		return
	}
	golog.Infof("GetModelEvaluationTask result: %s", result.UUID)
}

func TestUpdateModelEvaluationTask(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryModelEvaluationTask(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateModelEvaluationTask Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ModelEvaluationTask found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateModelEvaluationTaskParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateModelEvaluationTaskParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateModelEvaluationTask(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateModelEvaluationTask error: %v", err)
		return
	}
	golog.Infof("UpdateModelEvaluationTask result: %s", result.UUID)
}

func TestDeleteModelEvaluationTask(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteModelEvaluationTask is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryModelEvaluationTask(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteModelEvaluationTask Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ModelEvaluationTask found to test Delete")
		return
	}

	err = accountLoginCli.DeleteModelEvaluationTask(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteModelEvaluationTask error: %v", err)
		return
	}
	golog.Infof("DeleteModelEvaluationTask succeeded for UUID: %s", list[0].UUID)
}
