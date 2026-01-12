// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmNicSecurityPolicy(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVmNicSecurityPolicy(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmNicSecurityPolicy error: %v", err)
		return
	}
	golog.Infof("QueryVmNicSecurityPolicy result count: %d", len(result))
}
func TestGetVmNicSecurityPolicy(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVmNicSecurityPolicy(&queryParam)
	if err != nil {
		t.Errorf("TestGetVmNicSecurityPolicy Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VmNicSecurityPolicy found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVmNicSecurityPolicy(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVmNicSecurityPolicy error: %v", err)
		return
	}
	golog.Infof("GetVmNicSecurityPolicy result: %s", result.UUID)
}

func TestChangeVmNicSecurityPolicy(t *testing.T) {
	// Change operation
	t.Skip("TestChangeVmNicSecurityPolicy requires specific parameters")

}
