// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryIAM2OrganizationProjectRef queries IAM2OrganizationProjectRef list
func (cli *ZSClient) QueryIAM2OrganizationProjectRef(ctx context.Context, params *param.QueryParam) ([]view.IAM2OrganizationProjectRefInventoryView, error) {
	var resp []view.IAM2OrganizationProjectRefInventoryView
	return resp, cli.List(ctx, "v1/iam2/projects/organizations/refs", params, &resp)
}

func (cli *ZSClient) GetIAM2OrganizationProjectRef(ctx context.Context, uuid string) (*view.IAM2OrganizationProjectRefInventoryView, error) {
	var resp view.IAM2OrganizationProjectRefInventoryView
	if err := cli.Get(ctx, "v1/iam2/projects/organizations/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIAM2OrganizationProjectRef Pagination
func (cli *ZSClient) PageIAM2OrganizationProjectRef(ctx context.Context, params *param.QueryParam) ([]view.IAM2OrganizationProjectRefInventoryView, int, error) {
	var iAM2OrganizationProjectRefs []view.IAM2OrganizationProjectRefInventoryView
	total, err := cli.Page(ctx, "v1/iam2/projects/organizations/refs", params, &iAM2OrganizationProjectRefs)
	return iAM2OrganizationProjectRefs, total, err
}
