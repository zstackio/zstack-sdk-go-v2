// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryEventSubscription(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryEventSubscription(&queryParam)
	if err != nil {
		t.Errorf("TestQueryEventSubscription error: %v", err)
		return
	}
	golog.Infof("QueryEventSubscription result count: %d", len(result))
}
func TestGetEventSubscription(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryEventSubscription(&queryParam)
	if err != nil {
		t.Errorf("TestGetEventSubscription Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No EventSubscription found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetEventSubscription(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetEventSubscription error: %v", err)
		return
	}
	golog.Infof("GetEventSubscription result: %s", result.UUID)
}
