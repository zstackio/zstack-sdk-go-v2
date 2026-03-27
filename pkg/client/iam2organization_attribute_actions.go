// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateIAM2OrganizationAttribute updates IAM2OrganizationAttribute
func (cli *ZSClient) UpdateIAM2OrganizationAttribute(ctx context.Context, uuid string, params param.UpdateIAM2OrganizationAttributeParam) (*view.IAM2OrganizationAttributeInventoryView, error) {
	resp := view.IAM2OrganizationAttributeInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/organizations/attributes", uuid, "", map[string]interface{}{
		"updateIAM2OrganizationAttribute": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryIAM2OrganizationAttribute queries IAM2OrganizationAttribute list
func (cli *ZSClient) QueryIAM2OrganizationAttribute(ctx context.Context, params *param.QueryParam) ([]view.IAM2OrganizationAttributeInventoryView, error) {
	var resp []view.IAM2OrganizationAttributeInventoryView
	return resp, cli.List(ctx, "v1/iam2/organizations/attributes", params, &resp)
}

func (cli *ZSClient) GetIAM2OrganizationAttribute(ctx context.Context, uuid string) (*view.IAM2OrganizationAttributeInventoryView, error) {
	var resp view.IAM2OrganizationAttributeInventoryView
	if err := cli.Get(ctx, "v1/iam2/organizations/attributes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIAM2OrganizationAttribute Pagination
func (cli *ZSClient) PageIAM2OrganizationAttribute(ctx context.Context, params *param.QueryParam) ([]view.IAM2OrganizationAttributeInventoryView, int, error) {
	var iAM2OrganizationAttributes []view.IAM2OrganizationAttributeInventoryView
	total, err := cli.Page(ctx, "v1/iam2/organizations/attributes", params, &iAM2OrganizationAttributes)
	return iAM2OrganizationAttributes, total, err
}
