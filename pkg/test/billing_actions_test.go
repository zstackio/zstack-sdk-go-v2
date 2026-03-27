// Copyright (c) ZStack.io, Inc.

package test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestDeleteBilling(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteBilling is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAccountBilling(context.Background(), &queryParam)
	if err != nil {
		t.Errorf("TestDeleteBilling Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Billing found to test Delete")
		return
	}

	err = accountLoginCli.DeleteBilling(context.Background(), list[0].ResourceName, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteBilling error: %v", err)
		return
	}
	golog.Infof("DeleteBilling succeeded for UUID: %s", list[0].ResourceUuid)
}
