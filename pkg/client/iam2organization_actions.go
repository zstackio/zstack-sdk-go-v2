// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryIAM2Organization queries IAM2Organization list
func (cli *ZSClient) QueryIAM2Organization(params *param.QueryParam) ([]view.IAM2OrganizationInventoryView, error) {
	var resp []view.IAM2OrganizationInventoryView
	return resp, cli.List("v1/iam2/organizations", params, &resp)
}
// DeleteIAM2Organization deletes IAM2Organization
func (cli *ZSClient) DeleteIAM2Organization(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/organizations/{uuid}", uuid, string(deleteMode))
}
// UpdateIAM2Organization updates IAM2Organization
func (cli *ZSClient) UpdateIAM2Organization(uuid string, params param.UpdateIAM2OrganizationParam) (*view.IAM2OrganizationInventoryView, error) {
	var resp view.UpdateIAM2OrganizationEventView
	if err := cli.Put("v1/iam2/organizations/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateIAM2Organization creates IAM2Organization
func (cli *ZSClient) CreateIAM2Organization(params param.CreateIAM2OrganizationParam) (*view.IAM2OrganizationInventoryView, error) {
	var resp view.CreateIAM2OrganizationEventView
	if err := cli.Post("v1/iam2/organizations", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
