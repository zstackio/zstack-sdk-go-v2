// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHost(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryHost(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHost error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryHost result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s\t%s", r.UUID, r.Name, r.State, r.Status, r.HypervisorType)
	}
	golog.Infof("======================================")
}

func TestQueryHost2(t *testing.T) {
	// Query with conditions - Connected hosts
	params := param.NewQueryParam()
	params.AddQ("state=Enabled")
	params.AddQ("status=Connected")
	params.Start(0).Limit(10).ReplyWithCount(true)
	result, err := accessKeyAuthCli.QueryHost(&params)
	if err != nil {
		t.Errorf("TestQueryHost2 error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("Found %d Enabled/Connected Hosts:", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%d CPU\t%d MB RAM", r.UUID, r.Name, r.HypervisorType, r.TotalCpuCapacity, r.TotalMemoryCapacity/1024/1024)
	}
	golog.Infof("======================================")
}

func TestPageHost(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageHost(&queryParam)
	if err != nil {
		t.Errorf("TestPageHost error: %v", err)
		return
	}
	golog.Infof("PageHost result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s", r.UUID, r.Name, r.State, r.Status)
	}
}

func TestGetHost(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryHost(&queryParam)
	if err != nil {
		t.Errorf("TestGetHost Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Host found to test Get")
		return
	}

	// Get by UUID
	result, err := accessKeyAuthCli.GetHost(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetHost error: %v", err)
		return
	}
	golog.Infof("GetHost result: %s, Name: %s", result.UUID, result.Name)
}
