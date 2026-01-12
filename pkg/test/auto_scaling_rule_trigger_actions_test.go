// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAutoScalingRuleTrigger(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAutoScalingRuleTrigger(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAutoScalingRuleTrigger error: %v", err)
		return
	}
	golog.Infof("QueryAutoScalingRuleTrigger result count: %d", len(result))
}
func TestGetAutoScalingRuleTrigger(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAutoScalingRuleTrigger(&queryParam)
	if err != nil {
		t.Errorf("TestGetAutoScalingRuleTrigger Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AutoScalingRuleTrigger found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetAutoScalingRuleTrigger(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetAutoScalingRuleTrigger error: %v", err)
		return
	}
	golog.Infof("GetAutoScalingRuleTrigger result: %s", result.UUID)
}

func TestDeleteAutoScalingRuleTrigger(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteAutoScalingRuleTrigger is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAutoScalingRuleTrigger(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteAutoScalingRuleTrigger Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AutoScalingRuleTrigger found to test Delete")
		return
	}

	err = accountLoginCli.DeleteAutoScalingRuleTrigger(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteAutoScalingRuleTrigger error: %v", err)
		return
	}
	golog.Infof("DeleteAutoScalingRuleTrigger succeeded for UUID: %s", list[0].UUID)
}
