// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAccountPriceTableRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAccountPriceTableRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAccountPriceTableRef error: %v", err)
		return
	}
	golog.Infof("QueryAccountPriceTableRef result count: %d", len(result))
}

func TestGetAccountPriceTableRef(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAccountPriceTableRef(&queryParam)
	if err != nil {
		t.Errorf("TestGetAccountPriceTableRef Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AccountPriceTableRef found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetAccountPriceTableRef(list[0].AccountUuid)
	if err != nil {
		t.Errorf("TestGetAccountPriceTableRef error: %v", err)
		return
	}
	golog.Infof("GetAccountPriceTableRef result: %s", result.AccountUuid)
}
