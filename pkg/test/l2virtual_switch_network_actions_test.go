// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryL2VirtualSwitchNetwork(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryL2VirtualSwitchNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestQueryL2VirtualSwitchNetwork error: %v", err)
		return
	}
	golog.Infof("QueryL2VirtualSwitchNetwork result count: %d", len(result))
}
func TestGetL2VirtualSwitchNetwork(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryL2VirtualSwitchNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestGetL2VirtualSwitchNetwork Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No L2VirtualSwitchNetwork found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetL2VirtualSwitchNetwork(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetL2VirtualSwitchNetwork error: %v", err)
		return
	}
	golog.Infof("GetL2VirtualSwitchNetwork result: %s", result.UUID)
}
