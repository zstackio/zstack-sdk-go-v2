// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySftpBackupStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySftpBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySftpBackupStorage error: %v", err)
		return
	}
	golog.Infof("QuerySftpBackupStorage result count: %d", len(result))
}

