// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryL2PortGroupNetwork(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryL2PortGroupNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestQueryL2PortGroupNetwork error: %v", err)
		return
	}
	golog.Infof("QueryL2PortGroupNetwork result count: %d", len(result))
}

