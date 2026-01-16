// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVolumeSnapshot(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryVolumeSnapshot(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVolumeSnapshot error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryVolumeSnapshot result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s", r.UUID, r.Name, r.State, r.Type)
	}
	golog.Infof("======================================")
}

func TestPageVolumeSnapshot(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageVolumeSnapshot(&queryParam)
	if err != nil {
		t.Errorf("TestPageVolumeSnapshot error: %v", err)
		return
	}
	golog.Infof("PageVolumeSnapshot result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
}

func TestGetVolumeSnapshot(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryVolumeSnapshot(&queryParam)
	if err != nil {
		t.Errorf("TestGetVolumeSnapshot Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VolumeSnapshot found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetVolumeSnapshot(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVolumeSnapshot error: %v", err)
		return
	}
	golog.Infof("GetVolumeSnapshot result: %s, Name: %s", result.UUID, result.Name)
}
