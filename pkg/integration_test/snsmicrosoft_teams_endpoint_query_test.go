// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSMicrosoftTeamsEndpoint(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySNSMicrosoftTeamsEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSMicrosoftTeamsEndpoint error: %v", err)
		return
	}
	golog.Infof("QuerySNSMicrosoftTeamsEndpoint result count: %d", len(result))
}

