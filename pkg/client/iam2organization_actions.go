// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryIAM2Organization queries IAM2Organization list
func (cli *ZSClient) QueryIAM2Organization(params *param.QueryParam) ([]view.IAM2OrganizationInventoryView, error) {
	var resp []view.IAM2OrganizationInventoryView
	return resp, cli.List("v1/iam2/organizations", params, &resp)
}

// PageIAM2Organization Pagination
func (cli *ZSClient) PageIAM2Organization(params *param.QueryParam) ([]view.IAM2OrganizationInventoryView, int, error) {
	var iAM2Organizations []view.IAM2OrganizationInventoryView
	total, err := cli.Page("v1/iam2/organizations", params, &iAM2Organizations)
	return iAM2Organizations, total, err
}
// DeleteIAM2Organization deletes IAM2Organization
func (cli *ZSClient) DeleteIAM2Organization(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/organizations", uuid, string(deleteMode))
}
// UpdateIAM2Organization updates IAM2Organization
func (cli *ZSClient) UpdateIAM2Organization(uuid string, params param.UpdateIAM2OrganizationParam) (*view.IAM2OrganizationInventoryView, error) {
	resp := view.IAM2OrganizationInventoryView{}
	if err := cli.Put("v1/iam2/organizations", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateIAM2Organization creates IAM2Organization
func (cli *ZSClient) CreateIAM2Organization(params param.CreateIAM2OrganizationParam) (*view.IAM2OrganizationInventoryView, error) {
	resp := view.IAM2OrganizationInventoryView{}
	if err := cli.Post("v1/iam2/organizations", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
