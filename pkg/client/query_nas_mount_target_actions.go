// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryNasMountTarget queries NasMountTarget list
func (cli *ZSClient) QueryNasMountTarget(params *param.QueryParam) ([]view.NasMountTargetInventoryView, error) {
	var resp []view.NasMountTargetInventoryView
	return resp, cli.List("v1/primary-storage/nas/mount", params, &resp)
}
