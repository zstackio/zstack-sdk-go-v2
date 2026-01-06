// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateIAM2VirtualIDGroupAttribute updates IAM2VirtualIDGroupAttribute
func (cli *ZSClient) UpdateIAM2VirtualIDGroupAttribute(uuid string, params param.UpdateIAM2VirtualIDGroupAttributeParam) (*view.IAM2VirtualIDGroupAttributeInventoryView, error) {
	var resp view.UpdateIAM2VirtualIDGroupAttributeEventView
	if err := cli.Put("v1/iam2/projects/groups/attributes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryIAM2VirtualIDGroupAttribute queries IAM2VirtualIDGroupAttribute list
func (cli *ZSClient) QueryIAM2VirtualIDGroupAttribute(params *param.QueryParam) ([]view.IAM2VirtualIDGroupAttributeInventoryView, error) {
	var resp []view.IAM2VirtualIDGroupAttributeInventoryView
	return resp, cli.List("v1/iam2/projects/groups/attributes/", params, &resp)
}
