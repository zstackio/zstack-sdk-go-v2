// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCephBackupStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryCephBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCephBackupStorage error: %v", err)
		return
	}
	golog.Infof("QueryCephBackupStorage result count: %d", len(result))
}

func TestAddCephBackupStorage(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddCephBackupStorage requires valid creation parameters")

}
