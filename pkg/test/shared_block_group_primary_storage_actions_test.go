// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySharedBlockGroupPrimaryStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySharedBlockGroupPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySharedBlockGroupPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("QuerySharedBlockGroupPrimaryStorage result count: %d", len(result))
}

func TestAddSharedBlockGroupPrimaryStorage(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddSharedBlockGroupPrimaryStorage requires valid creation parameters")

}
