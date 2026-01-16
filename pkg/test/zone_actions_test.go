// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryZone(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accessKeyAuthCli.QueryZone(&queryParam)
	if err != nil {
		t.Errorf("TestQueryZone error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryZone result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s", r.UUID, r.Name, r.State, r.Description)
	}
	golog.Infof("======================================")
}

func TestPageZone(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageZone(&queryParam)
	if err != nil {
		t.Errorf("TestPageZone error: %v", err)
		return
	}
	golog.Infof("PageZone result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.State)
	}
}

func TestGetZone(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryZone(&queryParam)
	if err != nil {
		t.Errorf("TestGetZone Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Zone found to test Get")
		return
	}

	// Get by UUID
	result, err := accessKeyAuthCli.GetZone(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetZone error: %v", err)
		return
	}
	golog.Infof("GetZone result: %s, Name: %s", result.UUID, result.Name)
}
