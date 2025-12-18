// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryNasMountTarget queries NasMountTarget list
func (cli *ZSClient) QueryNasMountTarget(params param.QueryParam) ([]view.NasMountTargetInventoryView, error) {
	var resp []view.NasMountTargetInventoryView
	return resp, cli.List("v1/primary-storage/nas/mount", &params, &resp)
}
