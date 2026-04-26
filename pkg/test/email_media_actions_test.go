// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryEmailMedia(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryEmailMedia(&queryParam)
	if err != nil {
		t.Errorf("TestQueryEmailMedia error: %v", err)
		return
	}
	golog.Infof("QueryEmailMedia result count: %d", len(result))
}

