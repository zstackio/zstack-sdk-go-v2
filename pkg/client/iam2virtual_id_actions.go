// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryIAM2VirtualID queries IAM2VirtualID list
func (cli *ZSClient) QueryIAM2VirtualID(params *param.QueryParam) ([]view.IAM2VirtualIDInventoryView, error) {
	var resp []view.IAM2VirtualIDInventoryView
	return resp, cli.List("v1/iam2/virtual-ids", params, &resp)
}
// CreateIAM2VirtualID creates IAM2VirtualID
func (cli *ZSClient) CreateIAM2VirtualID(params param.CreateIAM2VirtualIDParam) (*view.IAM2VirtualIDInventoryView, error) {
	var resp view.CreateIAM2VirtualIDEventView
	if err := cli.Post("v1/iam2/virtual-ids", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteIAM2VirtualID deletes IAM2VirtualID
func (cli *ZSClient) DeleteIAM2VirtualID(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/virtual-ids/{uuid}", uuid, string(deleteMode))
}
// LoginIAM2VirtualID operates on IAM2VirtualID
func (cli *ZSClient) LoginIAM2VirtualID(uuid string, params param.LoginIAM2VirtualIDParam) (*view.SessionInventoryView, error) {
	var resp view.LoginIAM2VirtualIDView
	if err := cli.Put("v1/iam2/virtual-ids/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateIAM2VirtualID updates IAM2VirtualID
func (cli *ZSClient) UpdateIAM2VirtualID(uuid string, params param.UpdateIAM2VirtualIDParam) (*view.IAM2VirtualIDInventoryView, error) {
	var resp view.UpdateIAM2VirtualIDEventView
	if err := cli.Put("v1/iam2/virtual-ids/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
