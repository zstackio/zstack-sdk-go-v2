// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2OrganizationProjectRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryIAM2OrganizationProjectRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2OrganizationProjectRef error: %v", err)
		return
	}
	golog.Infof("QueryIAM2OrganizationProjectRef result count: %d", len(result))
}

