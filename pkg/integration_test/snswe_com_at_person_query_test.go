// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSWeComAtPerson(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySNSWeComAtPerson(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSWeComAtPerson error: %v", err)
		return
	}
	golog.Infof("QuerySNSWeComAtPerson result count: %d", len(result))
}

