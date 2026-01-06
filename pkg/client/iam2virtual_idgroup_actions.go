// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryIAM2VirtualIDGroup queries IAM2VirtualIDGroup list
func (cli *ZSClient) QueryIAM2VirtualIDGroup(params *param.QueryParam) ([]view.IAM2VirtualIDGroupInventoryView, error) {
	var resp []view.IAM2VirtualIDGroupInventoryView
	return resp, cli.List("v1/iam2/projects/groups", params, &resp)
}
// CreateIAM2VirtualIDGroup creates IAM2VirtualIDGroup
func (cli *ZSClient) CreateIAM2VirtualIDGroup(params param.CreateIAM2VirtualIDGroupParam) (*view.IAM2VirtualIDGroupInventoryView, error) {
	var resp view.CreateIAM2VirtualIDGroupEventView
	if err := cli.Post("v1/iam2/groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteIAM2VirtualIDGroup deletes IAM2VirtualIDGroup
func (cli *ZSClient) DeleteIAM2VirtualIDGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/groups/{uuid}", uuid, string(deleteMode))
}
// UpdateIAM2VirtualIDGroup updates IAM2VirtualIDGroup
func (cli *ZSClient) UpdateIAM2VirtualIDGroup(uuid string, params param.UpdateIAM2VirtualIDGroupParam) (*view.IAM2VirtualIDGroupInventoryView, error) {
	var resp view.UpdateIAM2VirtualIDGroupEventView
	if err := cli.Put("v1/iam2/projects/groups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
