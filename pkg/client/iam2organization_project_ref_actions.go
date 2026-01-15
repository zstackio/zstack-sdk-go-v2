// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryIAM2OrganizationProjectRef queries IAM2OrganizationProjectRef list
func (cli *ZSClient) QueryIAM2OrganizationProjectRef(params *param.QueryParam) ([]view.IAM2OrganizationProjectRefInventoryView, error) {
	var resp []view.IAM2OrganizationProjectRefInventoryView
	return resp, cli.List("v1/iam2/projects/organizations/refs", params, &resp)
}

// PageIAM2OrganizationProjectRef Pagination
func (cli *ZSClient) PageIAM2OrganizationProjectRef(params *param.QueryParam) ([]view.IAM2OrganizationProjectRefInventoryView, int, error) {
	var iAM2OrganizationProjectRefs []view.IAM2OrganizationProjectRefInventoryView
	total, err := cli.Page("v1/iam2/projects/organizations/refs", params, &iAM2OrganizationProjectRefs)
	return iAM2OrganizationProjectRefs, total, err
}
