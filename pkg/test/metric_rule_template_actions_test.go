// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMetricRuleTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMetricRuleTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMetricRuleTemplate error: %v", err)
		return
	}
	golog.Infof("QueryMetricRuleTemplate result count: %d", len(result))
}

func TestUpdateMetricRuleTemplate(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMetricRuleTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateMetricRuleTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MetricRuleTemplate found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateMetricRuleTemplateParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateMetricRuleTemplateParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateMetricRuleTemplate(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateMetricRuleTemplate error: %v", err)
		return
	}
	golog.Infof("UpdateMetricRuleTemplate result: %s", result.UUID)
}

func TestDeleteMetricRuleTemplate(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteMetricRuleTemplate is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMetricRuleTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteMetricRuleTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MetricRuleTemplate found to test Delete")
		return
	}

	err = accountLoginCli.DeleteMetricRuleTemplate(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteMetricRuleTemplate error: %v", err)
		return
	}
	golog.Infof("DeleteMetricRuleTemplate succeeded for UUID: %s", list[0].UUID)
}

func TestAddMetricRuleTemplate(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddMetricRuleTemplate requires valid creation parameters")

}
