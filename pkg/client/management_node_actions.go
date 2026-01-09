// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryManagementNode queries ManagementNode list
func (cli *ZSClient) QueryManagementNode(params *param.QueryParam) ([]view.ManagementNodeInventoryView, error) {
	var resp []view.ManagementNodeInventoryView
	return resp, cli.List("v1/management-nodes", params, &resp)
}

func (cli *ZSClient) GetManagementNode(uuid string) (*view.ManagementNodeInventoryView, error) {
	var resp view.ManagementNodeInventoryView
	if err := cli.Get("v1/management-nodes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
