// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryL2VlanNetwork(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryL2VlanNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestQueryL2VlanNetwork error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryL2VlanNetwork result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%d", r.UUID, r.Name, r.Vlan)
	}
	golog.Infof("======================================")
}

func TestPageL2VlanNetwork(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageL2VlanNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestPageL2VlanNetwork error: %v", err)
		return
	}
	golog.Infof("PageL2VlanNetwork result: total=%d, returned=%d", total, len(result))
}

func TestGetL2VlanNetwork(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryL2VlanNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestGetL2VlanNetwork Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No L2VlanNetwork found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetL2VlanNetwork(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetL2VlanNetwork error: %v", err)
		return
	}
	golog.Infof("GetL2VlanNetwork result: %s, Name: %s", result.UUID, result.Name)
}
