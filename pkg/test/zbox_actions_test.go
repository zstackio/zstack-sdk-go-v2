// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryZBox(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryZBox(&queryParam)
	if err != nil {
		t.Errorf("TestQueryZBox error: %v", err)
		return
	}
	golog.Infof("QueryZBox result count: %d", len(result))
}

func TestAddZBox(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddZBox requires valid creation parameters")

}
