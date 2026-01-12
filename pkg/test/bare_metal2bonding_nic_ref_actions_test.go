// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBareMetal2BondingNicRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBareMetal2BondingNicRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBareMetal2BondingNicRef error: %v", err)
		return
	}
	golog.Infof("QueryBareMetal2BondingNicRef result count: %d", len(result))
}
func TestGetBareMetal2BondingNicRef(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBareMetal2BondingNicRef(&queryParam)
	if err != nil {
		t.Errorf("TestGetBareMetal2BondingNicRef Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BareMetal2BondingNicRef found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetBareMetal2BondingNicRef(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetBareMetal2BondingNicRef error: %v", err)
		return
	}
	golog.Infof("GetBareMetal2BondingNicRef result: %s", result.UUID)
}
