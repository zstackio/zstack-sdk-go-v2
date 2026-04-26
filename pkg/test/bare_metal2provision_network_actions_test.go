// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBareMetal2ProvisionNetwork(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBareMetal2ProvisionNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBareMetal2ProvisionNetwork error: %v", err)
		return
	}
	golog.Infof("QueryBareMetal2ProvisionNetwork result count: %d", len(result))
}

