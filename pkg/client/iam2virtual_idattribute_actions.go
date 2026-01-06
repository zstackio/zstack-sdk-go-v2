// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateIAM2VirtualIDAttribute updates IAM2VirtualIDAttribute
func (cli *ZSClient) UpdateIAM2VirtualIDAttribute(uuid string, params param.UpdateIAM2VirtualIDAttributeParam) (*view.IAM2VirtualIDAttributeInventoryView, error) {
	var resp view.UpdateIAM2VirtualIDAttributeEventView
	if err := cli.Put("v1/iam2/virtual-ids/attributes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryIAM2VirtualIDAttribute queries IAM2VirtualIDAttribute list
func (cli *ZSClient) QueryIAM2VirtualIDAttribute(params *param.QueryParam) ([]view.IAM2VirtualIDAttributeInventoryView, error) {
	var resp []view.IAM2VirtualIDAttributeInventoryView
	return resp, cli.List("v1/iam2/virtual-ids/attributes", params, &resp)
}
