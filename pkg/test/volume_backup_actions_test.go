// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVolumeBackup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryVolumeBackup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVolumeBackup error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryVolumeBackup result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
	golog.Infof("======================================")
}

func TestPageVolumeBackup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageVolumeBackup(&queryParam)
	if err != nil {
		t.Errorf("TestPageVolumeBackup error: %v", err)
		return
	}
	golog.Infof("PageVolumeBackup result: total=%d, returned=%d", total, len(result))
}

func TestGetVolumeBackup(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVolumeBackup(&queryParam)
	if err != nil {
		t.Errorf("TestGetVolumeBackup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VolumeBackup found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetVolumeBackup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVolumeBackup error: %v", err)
		return
	}
	golog.Infof("GetVolumeBackup result: %s, Name: %s", result.UUID, result.Name)
}
