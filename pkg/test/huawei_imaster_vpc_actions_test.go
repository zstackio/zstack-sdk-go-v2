// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHuaweiIMasterVpc(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryHuaweiIMasterVpc(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHuaweiIMasterVpc error: %v", err)
		return
	}
	golog.Infof("QueryHuaweiIMasterVpc result count: %d", len(result))
}

func TestDeleteHuaweiIMasterVpc(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteHuaweiIMasterVpc is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryHuaweiIMasterVpc(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteHuaweiIMasterVpc Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No HuaweiIMasterVpc found to test Delete")
		return
	}

	err = accountLoginCli.DeleteHuaweiIMasterVpc(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteHuaweiIMasterVpc error: %v", err)
		return
	}
	golog.Infof("DeleteHuaweiIMasterVpc succeeded for UUID: %s", list[0].UUID)
}
