// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySecurityMachine queries SecurityMachine list
func (cli *ZSClient) QuerySecurityMachine(params *param.QueryParam) ([]view.SecurityMachineInventoryView, error) {
	var resp []view.SecurityMachineInventoryView
	return resp, cli.List("v1/security-machines", params, &resp)
}
