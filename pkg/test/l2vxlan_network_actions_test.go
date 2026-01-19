// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryL2VxlanNetwork(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryL2VxlanNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestQueryL2VxlanNetwork error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryL2VxlanNetwork result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%d", r.UUID, r.Name, r.Vni)
	}
	golog.Infof("======================================")
}

func TestPageL2VxlanNetwork(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageL2VxlanNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestPageL2VxlanNetwork error: %v", err)
		return
	}
	golog.Infof("PageL2VxlanNetwork result: total=%d, returned=%d", total, len(result))
}

func TestGetL2VxlanNetwork(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryL2VxlanNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestGetL2VxlanNetwork Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No L2VxlanNetwork found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetL2VxlanNetwork(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetL2VxlanNetwork error: %v", err)
		return
	}
	golog.Infof("GetL2VxlanNetwork result: %s, Name: %s", result.UUID, result.Name)
}
