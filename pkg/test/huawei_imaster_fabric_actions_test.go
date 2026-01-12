// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHuaweiIMasterFabric(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryHuaweiIMasterFabric(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHuaweiIMasterFabric error: %v", err)
		return
	}
	golog.Infof("QueryHuaweiIMasterFabric result count: %d", len(result))
}

func TestDeleteHuaweiIMasterFabric(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteHuaweiIMasterFabric is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryHuaweiIMasterFabric(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteHuaweiIMasterFabric Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No HuaweiIMasterFabric found to test Delete")
		return
	}

	err = accountLoginCli.DeleteHuaweiIMasterFabric(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteHuaweiIMasterFabric error: %v", err)
		return
	}
	golog.Infof("DeleteHuaweiIMasterFabric succeeded for UUID: %s", list[0].UUID)
}
