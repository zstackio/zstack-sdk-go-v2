// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCephPrimaryStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryCephPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCephPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("QueryCephPrimaryStorage result count: %d", len(result))
}

func TestAddCephPrimaryStorage(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddCephPrimaryStorage requires valid creation parameters")

}
