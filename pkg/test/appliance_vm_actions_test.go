// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryApplianceVm(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryApplianceVm(&queryParam)
	if err != nil {
		t.Errorf("TestQueryApplianceVm error: %v", err)
		return
	}
	golog.Infof("QueryApplianceVm result count: %d", len(result))
}
func TestGetApplianceVm(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryApplianceVm(&queryParam)
	if err != nil {
		t.Errorf("TestGetApplianceVm Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ApplianceVm found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetApplianceVm(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetApplianceVm error: %v", err)
		return
	}
	golog.Infof("GetApplianceVm result: %s", result.UUID)
}
