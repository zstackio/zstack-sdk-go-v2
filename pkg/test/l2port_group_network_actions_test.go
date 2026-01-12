// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryL2PortGroupNetwork(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryL2PortGroupNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestQueryL2PortGroupNetwork error: %v", err)
		return
	}
	golog.Infof("QueryL2PortGroupNetwork result count: %d", len(result))
}
func TestGetL2PortGroupNetwork(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryL2PortGroupNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestGetL2PortGroupNetwork Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No L2PortGroupNetwork found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetL2PortGroupNetwork(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetL2PortGroupNetwork error: %v", err)
		return
	}
	golog.Infof("GetL2PortGroupNetwork result: %s", result.UUID)
}
