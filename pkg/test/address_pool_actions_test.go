// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAddressPool(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAddressPool(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAddressPool error: %v", err)
		return
	}
	golog.Infof("QueryAddressPool result count: %d", len(result))
}
func TestGetAddressPool(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAddressPool(&queryParam)
	if err != nil {
		t.Errorf("TestGetAddressPool Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AddressPool found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetAddressPool(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetAddressPool error: %v", err)
		return
	}
	golog.Infof("GetAddressPool result: %s", result.UUID)
}
