// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHuaweiIMasterTenant(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryHuaweiIMasterTenant(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHuaweiIMasterTenant error: %v", err)
		return
	}
	golog.Infof("QueryHuaweiIMasterTenant result count: %d", len(result))
}

func TestDeleteHuaweiIMasterTenant(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteHuaweiIMasterTenant is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryHuaweiIMasterTenant(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteHuaweiIMasterTenant Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No HuaweiIMasterTenant found to test Delete")
		return
	}

	err = accountLoginCli.DeleteHuaweiIMasterTenant(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteHuaweiIMasterTenant error: %v", err)
		return
	}
	golog.Infof("DeleteHuaweiIMasterTenant succeeded for UUID: %s", list[0].UUID)
}
