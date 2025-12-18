// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetOrganizationQuotaUsage 获取OrganizationQuotaUsage详情
func (cli *ZSClient) GetOrganizationQuotaUsage(uuid string) (*view.GetOrganizationQuotaUsageView, error) {
	var resp view.GetOrganizationQuotaUsageView
	if err := cli.Get("v1/iam2/organizations/quota/{uuid}/usages", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

