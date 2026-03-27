// Copyright (c) ZStack.io, Inc.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateEventSubscriptionLabel(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryEventSubscription(context.Background(), &queryParam)
	if err != nil {
		t.Errorf("TestUpdateEventSubscriptionLabel Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No EventSubscriptionLabel found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateEventSubscriptionLabelParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateEventSubscriptionLabelParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateEventSubscriptionLabel(context.Background(), list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateEventSubscriptionLabel error: %v", err)
		return
	}
	golog.Infof("UpdateEventSubscriptionLabel result: %s", result.UUID)
}
