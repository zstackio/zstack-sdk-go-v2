// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryL3Network_V2(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryL3Network(&queryParam)
	if err != nil {
		t.Errorf("TestQueryL3Network error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryL3Network result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s", r.UUID, r.Name, r.State, r.Type)
	}
	golog.Infof("======================================")
}

func TestPageL3Network_V2(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageL3Network(&queryParam)
	if err != nil {
		t.Errorf("TestPageL3Network error: %v", err)
		return
	}
	golog.Infof("PageL3Network result: total=%d, returned=%d", total, len(result))
}

func TestGetL3Network_V2(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryL3Network(&queryParam)
	if err != nil {
		t.Errorf("TestGetL3Network Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No L3Network found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetL3Network(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetL3Network error: %v", err)
		return
	}
	golog.Infof("GetL3Network result: %s, Name: %s", result.UUID, result.Name)
}
