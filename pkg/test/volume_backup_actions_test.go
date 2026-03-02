// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVolumeBackup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVolumeBackup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVolumeBackup error: %v", err)
		return
	}
	golog.Infof("QueryVolumeBackup result count: %d", len(result))
}

