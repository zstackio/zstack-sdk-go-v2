// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryEventRuleTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryEventRuleTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryEventRuleTemplate error: %v", err)
		return
	}
	golog.Infof("QueryEventRuleTemplate result count: %d", len(result))
}

func TestUpdateEventRuleTemplate(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryEventRuleTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateEventRuleTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No EventRuleTemplate found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateEventRuleTemplateParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateEventRuleTemplateParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateEventRuleTemplate(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateEventRuleTemplate error: %v", err)
		return
	}
	golog.Infof("UpdateEventRuleTemplate result: %s", result.UUID)
}

func TestDeleteEventRuleTemplate(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteEventRuleTemplate is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryEventRuleTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteEventRuleTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No EventRuleTemplate found to test Delete")
		return
	}

	err = accountLoginCli.DeleteEventRuleTemplate(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteEventRuleTemplate error: %v", err)
		return
	}
	golog.Infof("DeleteEventRuleTemplate succeeded for UUID: %s", list[0].UUID)
}

func TestAddEventRuleTemplate(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddEventRuleTemplate requires valid creation parameters")

}
