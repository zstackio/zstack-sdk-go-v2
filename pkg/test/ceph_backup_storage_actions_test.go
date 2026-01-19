// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCephBackupStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryCephBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCephBackupStorage error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryCephBackupStorage result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
	golog.Infof("======================================")
}

func TestPageCephBackupStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageCephBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestPageCephBackupStorage error: %v", err)
		return
	}
	golog.Infof("PageCephBackupStorage result: total=%d, returned=%d", total, len(result))
}

func TestGetCephBackupStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryCephBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestGetCephBackupStorage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No CephBackupStorage found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetCephBackupStorage(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetCephBackupStorage error: %v", err)
		return
	}
	golog.Infof("GetCephBackupStorage result: %s, Name: %s", result.UUID, result.Name)
}
