// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestDeleteAutoScalingTemplate(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteAutoScalingTemplate is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAutoScalingVmTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteAutoScalingTemplate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AutoScalingTemplate found to test Delete")
		return
	}

	err = accountLoginCli.DeleteAutoScalingTemplate(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteAutoScalingTemplate error: %v", err)
		return
	}
	golog.Infof("DeleteAutoScalingTemplate succeeded for UUID: %s", list[0].UUID)
}
