// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryInstanceOffering(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryInstanceOffering(&queryParam)
	if err != nil {
		t.Errorf("TestQueryInstanceOffering error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryInstanceOffering result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%d CPU\t%d MB RAM", r.UUID, r.Name, r.CpuNum, r.MemorySize/1024/1024)
	}
	golog.Infof("======================================")
}

func TestQueryInstanceOffering2(t *testing.T) {
	// Query with conditions
	params := param.NewQueryParam()
	params.AddQ("state=Enabled")
	params.Start(0).Limit(10).ReplyWithCount(true)
	result, err := accessKeyAuthCli.QueryInstanceOffering(&params)
	if err != nil {
		t.Errorf("TestQueryInstanceOffering2 error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("Found %d Enabled InstanceOfferings:", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%d CPU\t%d MB RAM\t%s", r.UUID, r.Name, r.CpuNum, r.MemorySize/1024/1024, r.Type)
	}
	golog.Infof("======================================")
}

func TestPageInstanceOffering(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageInstanceOffering(&queryParam)
	if err != nil {
		t.Errorf("TestPageInstanceOffering error: %v", err)
		return
	}
	golog.Infof("PageInstanceOffering result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%d CPU\t%d MB RAM", r.UUID, r.Name, r.CpuNum, r.MemorySize/1024/1024)
	}
}

func TestGetInstanceOffering(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryInstanceOffering(&queryParam)
	if err != nil {
		t.Errorf("TestGetInstanceOffering Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No InstanceOffering found to test Get")
		return
	}

	// Get by UUID
	result, err := accessKeyAuthCli.GetInstanceOffering(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetInstanceOffering error: %v", err)
		return
	}
	golog.Infof("GetInstanceOffering result: %s, Name: %s, CPU: %d", result.UUID, result.Name, result.CpuNum)
}
