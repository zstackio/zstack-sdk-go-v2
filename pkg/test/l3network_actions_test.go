// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryL3Network(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryL3Network(&queryParam)
	if err != nil {
		t.Errorf("TestQueryL3Network error: %v", err)
		return
	}
	golog.Infof("QueryL3Network result count: %d", len(result))
}

