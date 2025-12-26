// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetOrganizationQuotaUsage gets OrganizationQuotaUsage by uuid
func (cli *ZSClient) GetOrganizationQuotaUsage(uuid string) (*view.GetOrganizationQuotaUsageView, error) {
	var resp view.GetOrganizationQuotaUsageView
	if err := cli.Get("v1/iam2/organizations/quota/{uuid}/usages", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
