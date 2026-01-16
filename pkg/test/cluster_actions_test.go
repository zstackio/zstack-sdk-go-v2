// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCluster(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryCluster(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCluster error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryCluster result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s", r.UUID, r.Name, r.State, r.HypervisorType)
	}
	golog.Infof("======================================")
}

func TestQueryCluster2(t *testing.T) {
	// Query with conditions
	params := param.NewQueryParam()
	params.AddQ("state=Enabled")
	params.Start(0).Limit(10).ReplyWithCount(true)
	result, err := accessKeyAuthCli.QueryCluster(&params)
	if err != nil {
		t.Errorf("TestQueryCluster2 error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("Found %d Enabled Clusters:", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s", r.UUID, r.Name, r.HypervisorType, r.ZoneUuid)
	}
	golog.Infof("======================================")
}

func TestPageCluster(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageCluster(&queryParam)
	if err != nil {
		t.Errorf("TestPageCluster error: %v", err)
		return
	}
	golog.Infof("PageCluster result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s", r.UUID, r.Name, r.State, r.HypervisorType)
	}
}

func TestGetCluster(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryCluster(&queryParam)
	if err != nil {
		t.Errorf("TestGetCluster Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Cluster found to test Get")
		return
	}

	// Get by UUID
	result, err := accessKeyAuthCli.GetCluster(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetCluster error: %v", err)
		return
	}
	golog.Infof("GetCluster result: %s, Name: %s", result.UUID, result.Name)
}
