// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBaremetalChassis(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBaremetalChassis(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBaremetalChassis error: %v", err)
		return
	}
	golog.Infof("QueryBaremetalChassis result count: %d", len(result))
}

