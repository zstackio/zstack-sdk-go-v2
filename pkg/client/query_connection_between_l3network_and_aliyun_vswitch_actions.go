// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryConnectionBetweenL3NetworkAndAliyunVSwitch queries ConnectionBetweenL3NetworkAndAliyunVSwitch list
func (cli *ZSClient) QueryConnectionBetweenL3NetworkAndAliyunVSwitch(params param.QueryParam) ([]view.ConnectionRelationShipInventoryView, error) {
	var resp []view.ConnectionRelationShipInventoryView
	return resp, cli.List("v1/hybrid/aliyun/relationships", &params, &resp)
}
