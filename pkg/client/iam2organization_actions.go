// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryIAM2Organization queries IAM2Organization list
func (cli *ZSClient) QueryIAM2Organization(ctx context.Context, params *param.QueryParam) ([]view.IAM2OrganizationInventoryView, error) {
	var resp []view.IAM2OrganizationInventoryView
	return resp, cli.List(ctx, "v1/iam2/organizations", params, &resp)
}

func (cli *ZSClient) GetIAM2Organization(ctx context.Context, uuid string) (*view.IAM2OrganizationInventoryView, error) {
	var resp view.IAM2OrganizationInventoryView
	if err := cli.Get(ctx, "v1/iam2/organizations", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIAM2Organization Pagination
func (cli *ZSClient) PageIAM2Organization(ctx context.Context, params *param.QueryParam) ([]view.IAM2OrganizationInventoryView, int, error) {
	var iAM2Organizations []view.IAM2OrganizationInventoryView
	total, err := cli.Page(ctx, "v1/iam2/organizations", params, &iAM2Organizations)
	return iAM2Organizations, total, err
}
// DeleteIAM2Organization deletes IAM2Organization
func (cli *ZSClient) DeleteIAM2Organization(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/iam2/organizations", uuid, string(deleteMode))
}
// UpdateIAM2Organization updates IAM2Organization
func (cli *ZSClient) UpdateIAM2Organization(ctx context.Context, uuid string, params param.UpdateIAM2OrganizationParam) (*view.IAM2OrganizationInventoryView, error) {
	resp := view.IAM2OrganizationInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/organizations", uuid, "", map[string]interface{}{
		"updateIAM2Organization": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateIAM2Organization creates IAM2Organization
func (cli *ZSClient) CreateIAM2Organization(ctx context.Context, params param.CreateIAM2OrganizationParam) (*view.IAM2OrganizationInventoryView, error) {
	resp := view.IAM2OrganizationInventoryView{}
	if err := cli.Post(ctx, "v1/iam2/organizations", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
