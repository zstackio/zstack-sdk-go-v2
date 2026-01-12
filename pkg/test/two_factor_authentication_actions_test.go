// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryTwoFactorAuthentication(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryTwoFactorAuthentication(&queryParam)
	if err != nil {
		t.Errorf("TestQueryTwoFactorAuthentication error: %v", err)
		return
	}
	golog.Infof("QueryTwoFactorAuthentication result count: %d", len(result))
}
func TestGetTwoFactorAuthentication(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryTwoFactorAuthentication(&queryParam)
	if err != nil {
		t.Errorf("TestGetTwoFactorAuthentication Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No TwoFactorAuthentication found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetTwoFactorAuthentication(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetTwoFactorAuthentication error: %v", err)
		return
	}
	golog.Infof("GetTwoFactorAuthentication result: %s", result.UUID)
}
