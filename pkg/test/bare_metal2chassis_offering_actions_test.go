// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBareMetal2ChassisOffering(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBareMetal2ChassisOffering(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBareMetal2ChassisOffering error: %v", err)
		return
	}
	golog.Infof("QueryBareMetal2ChassisOffering result count: %d", len(result))
}

