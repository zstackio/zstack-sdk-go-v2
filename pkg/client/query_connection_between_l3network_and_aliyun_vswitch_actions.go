// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryConnectionBetweenL3NetworkAndAliyunVSwitch queries ConnectionBetweenL3NetworkAndAliyunVSwitch list
func (cli *ZSClient) QueryConnectionBetweenL3NetworkAndAliyunVSwitch(params *param.QueryParam) ([]view.ConnectionRelationShipInventoryView, error) {
	var resp []view.ConnectionRelationShipInventoryView
	return resp, cli.List("v1/hybrid/aliyun/relationships", params, &resp)
}
