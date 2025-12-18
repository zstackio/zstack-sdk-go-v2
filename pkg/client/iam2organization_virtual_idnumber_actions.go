// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetIAM2OrganizationVirtualIDNumber 获取IAM2OrganizationVirtualIDNumber详情
func (cli *ZSClient) GetIAM2OrganizationVirtualIDNumber(uuid string) (*view.GetIAM2OrganizationVirtualIDNumberView, error) {
	var resp view.GetIAM2OrganizationVirtualIDNumberView
	if err := cli.Get("v1/iam2/organizations/{uuid}/virtualIDNumber", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

