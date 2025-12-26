// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateOrganizationQuota updates OrganizationQuota
func (cli *ZSClient) UpdateOrganizationQuota(uuid string, params param.UpdateOrganizationQuotaParam) (*view.UpdateOrganizationQuotaEventView, error) {
	resp := view.UpdateOrganizationQuotaEventView{}
	if err := cli.Put("v1/iam2/Organization/quotas/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
