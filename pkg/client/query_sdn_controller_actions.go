// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySdnController queries SdnController list
func (cli *ZSClient) QuerySdnController(params *param.QueryParam) ([]view.SdnControllerInventoryView, error) {
	var resp []view.SdnControllerInventoryView
	return resp, cli.List("v1/sdn-controllers", params, &resp)
}
