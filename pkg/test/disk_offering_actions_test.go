// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryDiskOffering(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryDiskOffering(&queryParam)
	if err != nil {
		t.Errorf("TestQueryDiskOffering error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryDiskOffering result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%d MB", r.UUID, r.Name, r.DiskSize/1024/1024)
	}
	golog.Infof("======================================")
}

func TestQueryDiskOffering2(t *testing.T) {
	// Query with conditions
	params := param.NewQueryParam()
	params.AddQ("state=Enabled")
	params.Start(0).Limit(10).ReplyWithCount(true)
	result, err := accessKeyAuthCli.QueryDiskOffering(&params)
	if err != nil {
		t.Errorf("TestQueryDiskOffering2 error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("Found %d Enabled DiskOfferings:", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%d GB\t%s", r.UUID, r.Name, r.DiskSize/1024/1024/1024, r.Type)
	}
	golog.Infof("======================================")
}

func TestPageDiskOffering(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageDiskOffering(&queryParam)
	if err != nil {
		t.Errorf("TestPageDiskOffering error: %v", err)
		return
	}
	golog.Infof("PageDiskOffering result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%d MB", r.UUID, r.Name, r.DiskSize/1024/1024)
	}
}

func TestGetDiskOffering(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryDiskOffering(&queryParam)
	if err != nil {
		t.Errorf("TestGetDiskOffering Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No DiskOffering found to test Get")
		return
	}

	// Get by UUID
	result, err := accessKeyAuthCli.GetDiskOffering(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetDiskOffering error: %v", err)
		return
	}
	golog.Infof("GetDiskOffering result: %s, Name: %s, Size: %d MB", result.UUID, result.Name, result.DiskSize/1024/1024)
}
