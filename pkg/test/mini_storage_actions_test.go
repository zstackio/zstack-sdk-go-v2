// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMiniStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMiniStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMiniStorage error: %v", err)
		return
	}
	golog.Infof("QueryMiniStorage result count: %d", len(result))
}

func TestAddMiniStorage(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddMiniStorage requires valid creation parameters")

}
