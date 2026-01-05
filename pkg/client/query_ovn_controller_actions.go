// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryOvnController queries OvnController list
func (cli *ZSClient) QueryOvnController(params *param.QueryParam) ([]view.OvnControllerInventoryView, error) {
	var resp []view.OvnControllerInventoryView
	return resp, cli.List("v1/ovn-controllers", params, &resp)
}
