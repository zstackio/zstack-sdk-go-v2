// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateOrganizationQuota 更新OrganizationQuota
func (cli *ZSClient) UpdateOrganizationQuota(uuid string, params param.UpdateOrganizationQuotaParam) (*view.UpdateOrganizationQuotaEventView, error) {
	resp := view.UpdateOrganizationQuotaEventView{}
	if err := cli.Put("v1/iam2/Organization/quotas/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

