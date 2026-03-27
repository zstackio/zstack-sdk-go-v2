// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSDingTalkAtPerson(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySNSDingTalkAtPerson(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSDingTalkAtPerson error: %v", err)
		return
	}
	golog.Infof("QuerySNSDingTalkAtPerson result count: %d", len(result))
}

