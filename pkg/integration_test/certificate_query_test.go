// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCertificate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryCertificate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCertificate error: %v", err)
		return
	}
	golog.Infof("QueryCertificate result count: %d", len(result))
}

