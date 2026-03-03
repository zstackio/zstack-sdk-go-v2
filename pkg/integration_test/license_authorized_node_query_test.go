// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryLicenseAuthorizedNode(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryLicenseAuthorizedNode(&queryParam)
	if err != nil {
		t.Errorf("TestQueryLicenseAuthorizedNode error: %v", err)
		return
	}
	golog.Infof("QueryLicenseAuthorizedNode result count: %d", len(result))
}

